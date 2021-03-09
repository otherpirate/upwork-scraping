package scrapping

import (
	"testing"

	"github.com/go-test/deep"
	"github.com/otherpirate/upwork-scraping/pkg/models"
	"github.com/otherpirate/upwork-scraping/pkg/queue/mock_queue"
	"github.com/otherpirate/upwork-scraping/pkg/services/mock_service"
	"github.com/otherpirate/upwork-scraping/pkg/settings"
	"github.com/otherpirate/upwork-scraping/pkg/store/mock_store"
)

func TestCrawler(t *testing.T) {
	settings.LoadConfigs()
	store := mock_store.NewStore()
	service, _ := mock_service.NewServicePath("../..")
	message := models.MessageUser{}
	queue, _ := mock_queue.NewQueue(message)
	upworkScrapping := NewUpWork(
		service,
		store,
		queue,
	)
	defer upworkScrapping.Finish()
	queue.Listening(upworkScrapping.Crawler)
	expectProfile := models.Profile{
		ID:          "~011d3adda6e4865468",
		Account:     "bobsuperworker",
		Employer:    "upwork",
		CreatedAt:   store.Profile.CreatedAt,
		UpdatedAt:   store.Profile.UpdatedAt,
		FirstName:   "Bob",
		LastName:    "SuperHardworker",
		FullName:    "Bob SuperHardworker",
		Email:       "b******t2@argyle.io",
		PhoneNumber: "+1 3478823951",
		PictureURL:  "./Profile_files/c1s5uKXKHAkmbAMlvDYlBVERPiRcVTlQN0Kck3OSeNjB6W69xAqG2pPHirqicJzqx-",
		Address: models.Address{
			Line1:      "555 U.S. 22",
			Line2:      "11",
			City:       "Bridgewater Township",
			State:      "NJ",
			PostalCode: "08807",
			Country:    "United States",
		},
		Employment: models.Employment{
			Status:       "active",
			JobTitle:     "Main DB architect | BestSoftware",
			HireDatetime: "1942-01-01T00:00:00Z",
		},
	}
	if diff := deep.Equal(expectProfile, store.Profile); diff != nil {
		t.Errorf("Profile\nExpect : %+v \nGot: %+v \nDifference: %+v", expectProfile, store.Profile, diff)
	}
}
