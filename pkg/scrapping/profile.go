package scrapping

import (
	"log"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/otherpirate/upwork-scraping/pkg/models"
)

func (u *Upwork) Profile() (models.Profile, error) {
	log.Println("Loading profile...")
	profile := models.Profile{}
	profile, err := u.loadProfileInfo(profile)
	if err != nil {
		return profile, err
	}
	profile, err = u.loadContactInfo(profile)
	if err != nil {
		return profile, err
	}
	log.Printf("Profiled loaded %s", profile.ID)
	return profile, nil
}

func (u *Upwork) loadProfileInfo(profile models.Profile) (models.Profile, error) {
	page, err := u.loadProfilePage()
	if err != nil {
		return profile, err
	}
	doc := soup.HTMLParse(page)
	publicProfiliLink := doc.Find("div", "role", "group").Find("a")
	profile.ID = cleanID(publicProfiliLink.Attrs()["href"])
	return profile, nil
}

func (u *Upwork) loadContactInfo(profile models.Profile) (models.Profile, error) {
	page, err := u.loadContactInfoPage()
	if err != nil {
		return profile, err
	}
	doc := soup.HTMLParse(page)
	userID := doc.Find("div", "data-test", "userId")
	name := doc.Find("div", "data-test", "userName")
	email := doc.Find("div", "data-test", "userEmail")
	phone := doc.Find("div", "data-test", "phone")
	addressLine1 := doc.Find("span", "data-test", "addressStreet")
	addressLine2 := doc.Find("span", "data-test", "addressStreet2")
	addressCity := doc.Find("span", "data-test", "addressCity")
	addressState := doc.Find("span", "data-test", "addressState")
	addressPostalCode := doc.Find("span", "data-test", "addressZip")
	addressCountry := doc.Find("span", "data-test", "addressCountry")
	profile.PlatformUserID = Clean(userID.Text())
	profile.SetNames(Clean(name.Text()))
	profile.Email = Clean(email.Text())
	profile.PhoneNumber = Clean(phone.Text())
	profile.Address = models.Address{
		Line1:      Clean(addressLine1.Text()),
		Line2:      Clean(addressLine2.Text()),
		City:       Clean(addressCity.Text()),
		State:      Clean(addressState.Text())[2:],
		PostalCode: Clean(addressPostalCode.Text()),
		Country:    Clean(addressCountry.Text()),
	}
	return profile, nil
}

func cleanID(url string) string {
	url = strings.ReplaceAll(url, "?viewMode=1", "")
	return strings.ReplaceAll(url, "https://www.upwork.com/freelancers/~", "")
}

func (u *Upwork) loadProfilePage() (string, error) {
	err := u.service.Navigate("https://www.upwork.com/freelancers")
	if err != nil {
		return "", err
	}
	return u.service.PageSource()
}

func (u *Upwork) loadContactInfoPage() (string, error) {
	err := u.service.Navigate("https://www.upwork.com/freelancers/settings/contactInfo")
	// TODO: Re-enter password
	// TODO: Click Edit button to get email info
	if err != nil {
		return "", err
	}
	return u.service.PageSource()
}
