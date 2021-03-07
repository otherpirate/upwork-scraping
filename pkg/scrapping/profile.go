package scrapping

import (
	"log"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/otherpirate/upwork-scraping/pkg/models"
)

func (u *Upwork) profile() (models.Profile, error) {
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
	pictureURL := doc.Find("img", "class", "up-avatar")
	profile.ID = cleanID(publicProfiliLink.Attrs()["href"])
	profile.PictureURL = pictureURL.Attrs()["src"]

	var jobDiv soup.Root
	for _, div := range doc.FindAll("div", "class", "up-card") {
		header := div.Find("h2")
		if header.Pointer == nil {
			continue
		}
		if strings.Contains(header.Text(), "Employment history") {
			jobDiv = div
			break
		}
	}
	if jobDiv.Pointer != nil {
		jobTitle := jobDiv.Find("h4")
		jobPeriod := jobDiv.FindAll("div")[10]
		periods := strings.Split(cleanString(jobPeriod.Text()), " - ")
		hireDatetime := periods[0]
		terminationDatetime := periods[1]
		status := "terminated"
		if terminationDatetime == "Present" {
			status = "active"
			terminationDatetime = ""
		}
		profile.Employment = models.Employment{
			Status:              status,
			JobTitle:            cleanString(jobTitle.Text()),
			HireDatetime:        formatDateTime(hireDatetime),
			TerminationDatetime: formatDateTime(terminationDatetime),
		}
	}
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
	profile.Account = cleanString(userID.Text())
	profile.Employer = "upwork"
	profile.SetNames(cleanString(name.Text()))
	profile.Email = cleanString(email.Text())
	profile.PhoneNumber = cleanString(phone.Text())
	profile.Address = models.Address{
		Line1:      cleanString(addressLine1.Text()),
		Line2:      cleanString(addressLine2.Text()),
		City:       cleanString(addressCity.Text()),
		State:      cleanString(addressState.Text())[2:],
		PostalCode: cleanString(addressPostalCode.Text()),
		Country:    cleanString(addressCountry.Text()),
	}
	return profile, nil
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
