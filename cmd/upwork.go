package main

import (
	"log"
	"os"

	"github.com/otherpirate/upwork-scraping/pkg/queue/rabbitmq_queue"
	"github.com/otherpirate/upwork-scraping/pkg/scrapping"
	"github.com/otherpirate/upwork-scraping/pkg/services/selenium_service"

	"github.com/otherpirate/upwork-scraping/pkg/settings"
	"github.com/otherpirate/upwork-scraping/pkg/store"
)

func main() {
	settings.LoadConfigs()
	store := store.StoreJSON{
		Path: settings.StorePath,
	}
	service, err := selenium_service.NewService()
	if err != nil {
		log.Printf("Could not start selenium service. Reason %v", err)
		os.Exit(1)
	}
	queue, err := rabbitmq_queue.NewRabbitQueue()
	if err != nil {
		log.Printf("Could not start queue service. Reason %v", err)
		os.Exit(1)
	}
	upworkScrapping := scrapping.NewUpWork(
		service,
		store,
		queue,
	)
	defer upworkScrapping.Finish()

	keepUp := make(chan bool)
	queue.Listening(upworkScrapping.Crawler)
	<-keepUp
	os.Exit(0)
}
