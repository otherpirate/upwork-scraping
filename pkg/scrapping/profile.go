package scrapping

import (
	"log"

	"github.com/anaskhan96/soup"
	"github.com/otherpirate/upwork-scraping/pkg/models"
)

func (u *Upwork) Profile() (models.Profile, error) {
	log.Println("Loading profile...")
	profile := models.Profile{}
	page, err := u.loadProfilePage()
	if err != nil {
		return profile, err
	}
	doc := soup.HTMLParse(page)
	publicProfiliLink := doc.Find("div", "role", "group").Find("a")
	profile.SetIDLink(publicProfiliLink.Attrs()["href"])
	log.Printf("Profiled loaded %s", profile.ID)
	return profile, nil
}

func (u *Upwork) loadProfilePage() (string, error) {
	err := u.service.Navigate("https://www.upwork.com/freelancers")
	if err != nil {
		return "", err
	}
	return u.service.PageSource()
}
