package main

import (
	"log"
	"os"

	"github.com/gosimple/slug"
	"github.com/otherpirate/upwork-scraping/pkg/scrapping"
	"github.com/otherpirate/upwork-scraping/pkg/settings"
	"github.com/otherpirate/upwork-scraping/pkg/store"
)

func main() {
	settings.LoadConfigs()
	upworkScrapping, err := scrapping.NewUpWork(settings.UserName, settings.Password, settings.SecretAwnser)
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

	jobs, err := upworkScrapping.Jobs()
	if err != nil {
		log.Printf("Could not load jobs. Reason %v", err)
		os.Exit(1)
	}

	store := store.StoreJSON{
		Path: settings.StorePath,
	}
	for _, job := range jobs {
		name := slug.Make(job.Title)
		err = store.Save(name, job)
		if err != nil {
			log.Printf("Could not save jobs. Reason %v", err)
			os.Exit(1)
		}
	}
	os.Exit(0)
}
