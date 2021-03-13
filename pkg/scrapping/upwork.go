package scrapping

import (
	"log"

	"github.com/gosimple/slug"
	"github.com/otherpirate/upwork-scraping/pkg/models"
	"github.com/otherpirate/upwork-scraping/pkg/queue"
	"github.com/otherpirate/upwork-scraping/pkg/services"
	"github.com/otherpirate/upwork-scraping/pkg/store"
)

type Upwork struct {
	service services.Service
	store   store.Store
	queue   queue.Queue
}

func NewUpWork(service services.Service, store store.Store, queue queue.Queue) *Upwork {
	return &Upwork{
		service: service,
		store:   store,
		queue:   queue,
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

func (u *Upwork) Crawler(message models.MessageUser) error {
	u.service.Clear()
	err := u.login(message.UserName, message.Password, message.SecretAwnser)
	if err != nil {
		log.Printf("Could not login into Upwork. Reason %v", err)
		return err
	}

	jobs, err := u.jobs()
	if err != nil {
		log.Printf("Could not load jobs. Reason %v", err)
		return err
	}
	for _, job := range jobs {
		name := slug.Make(job.Title)
		err = u.store.SaveJob(name, &job)
		if err != nil {
			log.Printf("Could not save jobs. Reason %v", err)
			return err
		}
	}

	profile, err := u.profile(message.Password, message.ProfileData)
	if err != nil {
		log.Printf("Could not load profile. Reason %v", err)
		return err
	}
	err = u.store.SaveProfile(&profile)
	if err != nil {
		log.Printf("Could not save profile. Reason %v", err)
		return err
	}
	return u.queue.Foward(profile)
}
