package store

import "github.com/otherpirate/upwork-scraping/pkg/models"

type Store interface {
	SaveProfile(profile *models.Profile) error
	SaveJob(name string, job *models.Job) error
}
