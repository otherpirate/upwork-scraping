package mock_store

import (
	"time"

	"github.com/otherpirate/upwork-scraping/pkg/models"
)

const fileModePerm = 0755

func NewMockStore() mockStore {
	return mockStore{
		profiles: map[string]models.Profile{},
		jobs:     []models.Job{},
	}
}

type mockStore struct {
	profiles map[string]models.Profile
	jobs     []models.Job
}

func (s *mockStore) SaveProfile(profile *models.Profile) error {
	profile.CreatedAt = time.Now().UTC().Format("2006-01-02T15:04:05.999999Z")
	profile.UpdatedAt = profile.CreatedAt
	loadedProfile, exists := s.profiles[profile.ID]
	if exists {
		profile.CreatedAt = loadedProfile.CreatedAt
	}
	s.profiles[profile.ID] = *profile
	return nil
}

func (s *mockStore) SaveJob(name string, job models.Job) error {
	s.jobs = append(s.jobs, job)
	return nil
}
