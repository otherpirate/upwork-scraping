package main

import (
	"log"
	"os"

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
	upworkScrapping := scrapping.NewUpWork(
		service,
	)
	defer upworkScrapping.Finish()
	upworkScrapping.Crawler(store, settings.UserName, settings.Password, settings.SecretAwnser)
	os.Exit(0)
}
