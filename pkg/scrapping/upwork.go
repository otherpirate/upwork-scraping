package scrapping

import (
	"log"

	"github.com/gosimple/slug"
	"github.com/otherpirate/upwork-scraping/pkg/services"
	"github.com/otherpirate/upwork-scraping/pkg/store"
)

type Upwork struct {
	service services.Service
}

func NewUpWork(service services.Service) *Upwork {
	return &Upwork{
		service: service,
	}
}

func (u *Upwork) Finish() {
	if u == nil {
		return
	}
	if u.service != nil {
		u.service.Close()
	}
}

func (u *Upwork) Crawler(store store.StoreJSON, userName, password, secretAwnser string) {
	err := u.login(userName, password, secretAwnser)
	if err != nil {
		log.Printf("Could not login into Upwork. Reason %v", err)
		return
	}

	profile, err := u.profile()
	if err != nil {
		log.Printf("Could not load profile. Reason %v", err)
		return
	}
	err = store.SaveProfile(profile)
	if err != nil {
		log.Printf("Could not save profile. Reason %v", err)
		return
	}

	jobs, err := u.jobs()
	if err != nil {
		log.Printf("Could not load jobs. Reason %v", err)
		return
	}

	for _, job := range jobs {
		name := slug.Make(job.Title)
		err = store.SaveJob(name, job)
		if err != nil {
			log.Printf("Could not save jobs. Reason %v", err)
			return
		}
	}
}
