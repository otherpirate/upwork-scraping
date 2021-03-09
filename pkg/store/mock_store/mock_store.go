package mock_store

import (
	"time"

	"github.com/otherpirate/upwork-scraping/pkg/models"
)

const fileModePerm = 0755

func NewStore() *mockStore {
	return &mockStore{
		Profile: models.Profile{},
		Jobs:    []models.Job{},
	}
}

type mockStore struct {
	Profile models.Profile
	Jobs    []models.Job
}

func (s *mockStore) SaveProfile(profile *models.Profile) error {
	profile.CreatedAt = time.Now().UTC().Format("2006-01-02T15:04:05.999999Z")
	profile.UpdatedAt = profile.CreatedAt
	loadedProfile := s.Profile
	if loadedProfile.CreatedAt != "" {
		profile.CreatedAt = loadedProfile.CreatedAt
	}
	s.Profile = *profile
	return nil
}

func (s *mockStore) SaveJob(name string, job *models.Job) error {
	s.Jobs = append(s.Jobs, *job)
	return nil
}
