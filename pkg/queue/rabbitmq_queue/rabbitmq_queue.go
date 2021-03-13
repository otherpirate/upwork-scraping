package rabbitmq_queue

import (
	"encoding/json"
	"log"
	"time"

	"github.com/otherpirate/upwork-scraping/pkg/models"
	"github.com/otherpirate/upwork-scraping/pkg/settings"
	"github.com/otherpirate/upwork-scraping/pkg/utils"
	"github.com/streadway/amqp"
)

type rabbitQueue struct {
	channelProcess *amqp.Channel
	channelRequeue *amqp.Channel
	queueUser      amqp.Queue
	queueProfile   amqp.Queue
	queueRetry     amqp.Queue
	queueFailure   amqp.Queue
}

func NewRabbitQueue() (*rabbitQueue, error) {
	conn, err := amqp.Dial(settings.RabbitURI)
	if err != nil {
		return &rabbitQueue{}, err
	}
	channelProcess, err := conn.Channel()
	if err != nil {
		return &rabbitQueue{}, err
	}
	channelRequeue, err := conn.Channel()
	if err != nil {
		return &rabbitQueue{}, err
	}
	queueUser, err := delcareQueue(settings.RabbitQueueUser, channelProcess)
	if err != nil {
		return &rabbitQueue{}, err
	}
	queueProfile, err := delcareQueue(settings.RabbitQueueProfile, channelProcess)
	if err != nil {
		return &rabbitQueue{}, err
	}
	queueRetry, err := delcareQueue(settings.RabbitQueueRetry, channelRequeue)
	if err != nil {
		return &rabbitQueue{}, err
	}
	queueFailure, err := delcareQueue(settings.RabbitQueueFailure, channelRequeue)
	if err != nil {
		return &rabbitQueue{}, err
	}
	return &rabbitQueue{
		channelProcess: channelProcess,
		channelRequeue: channelRequeue,
		queueUser:      queueUser,
		queueProfile:   queueProfile,
		queueRetry:     queueRetry,
		queueFailure:   queueFailure,
	}, nil
}

func delcareQueue(name string, channel *amqp.Channel) (amqp.Queue, error) {
	return channel.QueueDeclare(
		name,
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

func (q *rabbitQueue) Listening(crawler func(message models.MessageUser) error) error {
	err := q.listening(crawler)
	if err != nil {
		return err
	}
	err = q.requeueFailed()
	if err != nil {
		return err
	}
	return nil
}

func (q *rabbitQueue) listening(crawler func(message models.MessageUser) error) error {
	msgs, err := consumer(q.queueUser.Name, "upwork-scraping", q.channelProcess)
	if err != nil {
		return err
	}
	go func() {
		for msg := range msgs {
			message := models.MessageUser{}
			json.Unmarshal(msg.Body, &message)
			log.Printf("Received a message: %+v", message)
			err := crawler(message)
			if err != nil {
				log.Printf("Could not process message: Reason %v", err)
				q.putInRetryQueue(message)
			}
			msg.Ack(false)
		}
	}()
	return nil
}

func (q *rabbitQueue) putInRetryQueue(message models.MessageUser) {
	message.Retries++
	err := publisher(q.queueRetry.Name, message, q.channelProcess)
	if err != nil {
		log.Println("Could not requeue the message because:", err)
	}
}

func (q *rabbitQueue) requeueFailed() error {
	msgs, err := consumer(q.queueRetry.Name, "upwork-scraping", q.channelRequeue)
	if err != nil {
		return err
	}
	go func() {
		for msg := range msgs {
			time.Sleep(time.Duration(settings.WaitBeforeRequeue))
			message := models.MessageUser{}
			json.Unmarshal(msg.Body, &message)
			if message.Retries > settings.CrawlerMaxRetries {
				log.Printf("Message %+v have to many retries. Put it a long term queue", message)
				publisher(q.queueFailure.Name, message, q.channelRequeue)
			} else {
				log.Printf("Requeue a message: %+v", message)
				publisher(q.queueUser.Name, message, q.channelRequeue)
			}
			msg.Ack(false)
		}

	}()
	return nil
}

func (q *rabbitQueue) Foward(profile models.Profile) error {
	return publisher(q.queueProfile.Name, profile, q.channelProcess)
}

func consumer(name, consumer string, channel *amqp.Channel) (<-chan amqp.Delivery, error) {
	return channel.Consume(
		name,     // queue
		consumer, // consumer
		false,    // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
}

func publisher(routingKey string, value interface{}, channel *amqp.Channel) error {
	json, err := utils.ToJSON(value)
	if err != nil {
		return err
	}
	err = channel.Publish(
		"",         // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        json,
		},
	)
	return err
}
