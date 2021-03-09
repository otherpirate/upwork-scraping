package mock_queue

import (
	"github.com/otherpirate/upwork-scraping/pkg/models"
)

type mockQueue struct {
	message       models.MessageUser
	FowardProfile models.Profile
}

func NewQueue(message models.MessageUser) (*mockQueue, error) {
	return &mockQueue{
		message: message,
	}, nil
}

func (q *mockQueue) Listening(crawler func(message models.MessageUser) error) error {
	return crawler(q.message)
}

func (q *mockQueue) Foward(profile models.Profile) error {
	q.FowardProfile = profile
	return nil
}
