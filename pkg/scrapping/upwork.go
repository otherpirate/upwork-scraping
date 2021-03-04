package scrapping

import "github.com/otherpirate/upwork-scraping/pkg/utils"

const loginUrl = "https://www.upwork.com/ab/account-security/login"

type Upwork struct {
	userName     string
	password     string
	secretAwnser string
	loginUrl     string
	service      *utils.SeleniumService
}

func NewUpWork(userName, password, secretAwnser string) (*Upwork, error) {
	service, err := utils.NewSeleniumService()
	if err != nil {
		return nil, err
	}
	return &Upwork{
		userName:     userName,
		password:     password,
		secretAwnser: secretAwnser,
		loginUrl:     loginUrl,
		service:      service,
	}, nil
}

func (u *Upwork) Finish() {
	if u == nil {
		return
	}
	if u.service != nil {
		u.service.Close()
	}
}
