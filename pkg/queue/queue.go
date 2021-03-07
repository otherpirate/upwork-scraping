package queue

import "github.com/otherpirate/upwork-scraping/pkg/models"

type Queue interface {
	Listening(crawler func(userName, password, secretAwnser string) error) error
	Foward(models.Profile) error
}
