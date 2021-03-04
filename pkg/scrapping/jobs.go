package scrapping

import (
	"time"

	"github.com/anaskhan96/soup"
	"github.com/otherpirate/upwork-scraping/pkg/models"
)

func (u *Upwork) Jobs() ([]models.Job, error) {
	jobs := []models.Job{}
	page, err := u.loadJobPage()
	if err != nil {
		return jobs, err
	}
	doc := soup.HTMLParse(page)
	for _, jobSection := range doc.FindAll("section", "class", "job-tile") {
		title := jobSection.Find("h4", "class", "job-title").Find("a")
		descriptionSection := jobSection.Find("div", "class", "description")
		resume := descriptionSection.Find("span", "class", "ng-binding")
		job := models.Job{
			Title:  title.Text(),
			Link:   title.Attrs()["href"],
			Resume: Clean(resume.Text()),
		}
		jobs = append(jobs, job)

	}
	return jobs, nil
}

func (u *Upwork) loadJobPage() (string, error) {
	err := u.service.Navigate("https://www.upwork.com/ab/find-work/domestic")
	if err != nil {
		return "", err
	}
	time.Sleep(stepWait)
	return u.service.WebDriver.PageSource()
}
