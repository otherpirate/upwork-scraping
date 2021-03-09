package queue

import "github.com/otherpirate/upwork-scraping/pkg/models"

type Queue interface {
	Listening(crawler func(models.MessageUser) error) error
	Foward(models.Profile) error
}
