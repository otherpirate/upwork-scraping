package scrapping

import (
	"fmt"
	"log"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/otherpirate/upwork-scraping/pkg/models"
)

func (u *Upwork) profile(password string, profile models.Profile) (models.Profile, error) {
	log.Println("Loading profile...")
	err := u.reenterPassword(password)
	if err != nil {
		return profile, err
	}
	profile, err = u.loadContactInfo(profile)
	if err != nil {
		return profile, err
	}
	profile, err = u.loadProfileInfo(profile)
	if err != nil {
		return profile, err
	}
	log.Printf("Profiled loaded %s", profile.ID)
	return profile, nil
}

func (u *Upwork) loadProfileInfo(profile models.Profile) (models.Profile, error) {
	page, err := u.loadProfilePage(profile.ID)
	if err != nil {
		return profile, err
	}
	doc := soup.HTMLParse(page)
	pictureURL := doc.Find("img", "class", "up-avatar")
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
	subMenuLinks := doc.FindAll("a", "data-cy", "menu-item-trigger")
	userID := doc.Find("div", "data-test", "userId")
	name := doc.Find("div", "data-test", "userName")
	// TODO: Click Edit button to get email info
	email := doc.Find("div", "data-test", "userEmail")
	phone := doc.Find("div", "data-test", "phone")
	addressLine1 := doc.Find("span", "data-test", "addressStreet")
	addressLine2 := doc.Find("span", "data-test", "addressStreet2")
	addressCity := doc.Find("span", "data-test", "addressCity")
	addressState := doc.Find("span", "data-test", "addressState")
	addressPostalCode := doc.Find("span", "data-test", "addressZip")
	addressCountry := doc.Find("span", "data-test", "addressCountry")
	profile.ID = cleanID(subMenuLinks)
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

func (u *Upwork) loadProfilePage(profileID string) (string, error) {
	err := u.service.Navigate(fmt.Sprintf("https://www.upwork.com/freelancers/%s", profileID))
	if err != nil {
		return "", err
	}
	return u.service.PageSource()
}

func (u *Upwork) loadContactInfoPage() (string, error) {
	err := u.service.Navigate("https://www.upwork.com/freelancers/settings/contactInfo")
	if err != nil {
		return "", err
	}
	return u.service.PageSource()
}

func (u Upwork) reenterPassword(password string) error {
	err := u.service.Navigate("https://www.upwork.com/ab/account-security/reenter-password")
	if err != nil {
		return err
	}
	passInput, err := u.service.WaitElement("id", "sensitiveZone_password")
	if err != nil {
		return err
	}
	passInput.SendKeys(password)
	continueButton, err := u.service.WaitElement("id", "control_continue")
	if err != nil {
		return err
	}
	return continueButton.Click()
}
