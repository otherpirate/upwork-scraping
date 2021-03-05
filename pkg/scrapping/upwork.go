package scrapping

import (
	"github.com/otherpirate/upwork-scraping/pkg/services"
)

type Upwork struct {
	userName     string
	password     string
	secretAwnser string
	service      services.Service
}

func NewUpWork(userName, password, secretAwnser string, service services.Service) *Upwork {
	return &Upwork{
		userName:     userName,
		password:     password,
		secretAwnser: secretAwnser,
		service:      service,
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
