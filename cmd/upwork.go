package main

import (
	"log"
	"os"

	"github.com/gosimple/slug"
	"github.com/otherpirate/upwork-scraping/pkg/scrapping"
	"github.com/otherpirate/upwork-scraping/pkg/services/selenium_service"

	//"github.com/otherpirate/upwork-scraping/pkg/services/selenium_service"

	"github.com/otherpirate/upwork-scraping/pkg/settings"
	"github.com/otherpirate/upwork-scraping/pkg/store"
)

func main() {
	settings.LoadConfigs()
	store := store.StoreJSON{
		Path: settings.StorePath,
	}
	service, err := selenium_service.NewService()
	//service, err := mock_service.NewService()
	if err != nil {
		log.Printf("Could not start selenium service. Reason %v", err)
		os.Exit(1)
	}
	upworkScrapping := scrapping.NewUpWork(
		settings.UserName,
		settings.Password,
		settings.SecretAwnser,
		service,
	)
	defer upworkScrapping.Finish()
	if err != nil {
		log.Printf("Could not start Upwork scrapping. Reason %v", err)
		os.Exit(1)
	}

	err = upworkScrapping.Login()
	if err != nil {
		log.Printf("Could not login into Upwork. Reason %v", err)
		os.Exit(1)
	}

	profile, err := upworkScrapping.Profile()
	if err != nil {
		log.Printf("Could not load profile. Reason %v", err)
		os.Exit(1)
	}
	err = store.SaveProfile(profile)
	if err != nil {
		log.Printf("Could not save profile. Reason %v", err)
		os.Exit(1)
	}

	jobs, err := upworkScrapping.Jobs()
	if err != nil {
		log.Printf("Could not load jobs. Reason %v", err)
		os.Exit(1)
	}

	for _, job := range jobs {
		name := slug.Make(job.Title)
		err = store.SaveJob(name, job)
		if err != nil {
			log.Printf("Could not save jobs. Reason %v", err)
			os.Exit(1)
		}
	}
	os.Exit(0)
}
