package rabbitmq_queue

import (
	"encoding/json"
	"log"

	"github.com/otherpirate/upwork-scraping/pkg/models"
	"github.com/otherpirate/upwork-scraping/pkg/settings"
	"github.com/otherpirate/upwork-scraping/pkg/utils"
	"github.com/streadway/amqp"
)

type RabbitQueue struct {
	channel      *amqp.Channel
	queueUser    amqp.Queue
	queueProfile amqp.Queue
}

func NewRabbitQueue() (*RabbitQueue, error) {
	conn, err := amqp.Dial(settings.RabbitURI)
	if err != nil {
		return &RabbitQueue{}, err
	}
	channel, err := conn.Channel()
	if err != nil {
		return &RabbitQueue{}, err
	}
	queueUser, err := channel.QueueDeclare(
		settings.RabbitQueueUser,
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	queueProfile, err := channel.QueueDeclare(
		settings.RabbitQueueProfile,
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	return &RabbitQueue{
		channel:      channel,
		queueUser:    queueUser,
		queueProfile: queueProfile,
	}, nil
}

func (q *RabbitQueue) Listening(crawler func(userName, password, secretAwnser string) error) error {
	msgs, err := q.channel.Consume(
		q.queueUser.Name,  // queue
		"upwork-scraping", // consumer
		true,              // auto-ack
		true,              // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	if err != nil {
		return err
	}
	go func() {
		for msg := range msgs {
			message := models.MessageUser{}
			json.Unmarshal(msg.Body, &message)
			log.Printf("Received a message: %+v", message)
			err := crawler(message.UserName, message.Password, message.SecretAwnser)
			if err != nil {
				log.Printf("Could not process message: Reason %v", err)
				msg.Nack(false, true)
				continue
			}
			msg.Ack(false)
		}
	}()
	return nil
}

func (q *RabbitQueue) Foward(profile models.Profile) error {
	json, err := utils.ToJSON(profile)
	if err != nil {
		return err
	}
	err = q.channel.Publish(
		"",                  // exchange
		q.queueProfile.Name, // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        json,
		},
	)
	return err
}
