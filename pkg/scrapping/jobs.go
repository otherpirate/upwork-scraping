package scrapping

import (
	"log"
	"strings"

	"github.com/anaskhan96/soup"
	"github.com/otherpirate/upwork-scraping/pkg/models"
)

func (u *Upwork) jobs() ([]models.Job, error) {
	log.Println("Loading jobs...")
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
		jobType := jobSection.Find("strong", "class", "js-type")
		postedAt := jobSection.Find("span", "class", "js-posted").Find("time")
		proposals := jobSection.Find("span", "data-job-proposals", "jsuJobTileMediumController.job.proposalsTier").Find("strong")
		paymentVerified := jobSection.Find("span", "class", "payment-status")
		spent := jobSection.Find("span", "class", "client-spendings").Find("strong")
		location := jobSection.Find("strong", "class", "client-location")
		job := models.Job{
			Title:           title.Text(),
			Link:            title.Attrs()["href"],
			Resume:          cleanString(resume.Text()),
			Type:            cleanString(jobType.Text()),
			PostedAt:        postedAt.Attrs()["datetime"],
			Proposals:       cleanString(proposals.Text()),
			PaymentVerified: paymentVerified.Pointer != nil,
			Spent:           cleanString(spent.Text()),
			Location:        cleanString(location.Text()),
		}
		level := jobSection.Find("span", "class", "js-contractor-tier")
		if level.Pointer != nil {
			levelStr := strings.Replace(level.Text(), " - ", "", 1)
			job.Level = cleanString(levelStr)
		}
		estimate := jobSection.Find("span", "class", "js-duration")
		if estimate.Pointer != nil {
			job.Estimate = cleanString(estimate.Text())
		}
		for _, skill := range jobSection.FindAll("a", "class", "o-tag-skill") {
			job.Skills = append(job.Skills, cleanString(skill.Find("span").Text()))
		}
		jobs = append(jobs, job)
	}
	log.Printf("Jobs loaded %d", len(jobs))
	return jobs, nil
}

func (u *Upwork) loadJobPage() (string, error) {
	err := u.service.Navigate("https://www.upwork.com/ab/find-work/domestic")
	if err != nil {
		return "", err
	}
	return u.service.PageSource()
}
