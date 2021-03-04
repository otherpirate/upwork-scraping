package main

import (
	"log"
	"os"

	"github.com/otherpirate/upwork-scraping/pkg/scrapping"
	"github.com/otherpirate/upwork-scraping/pkg/settings"
)

func main() {
	settings.LoadConfigs()
	upworkScrapping, err := scrapping.NewUpWork(settings.UserName, settings.Password, settings.SecretAwnser)
	if err != nil {
		upworkScrapping.Finish()
		log.Printf("Could not start Upwork scrapping. Reason %v", err)
		os.Exit(1)
	}

	err = upworkScrapping.Login()
	if err != nil {
		upworkScrapping.Finish()
		log.Printf("Could not login into Upwork. Reason %v", err)
		os.Exit(1)
	}

	jobs, err := upworkScrapping.Jobs()
	if err != nil {
		upworkScrapping.Finish()
		log.Printf("Could not load jobs. Reason %v", err)
		os.Exit(1)
	}
	for _, job := range jobs {
		log.Println(job)
	}
	upworkScrapping.Finish()
	os.Exit(0)
}
