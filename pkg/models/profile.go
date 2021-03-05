package models

import "strings"

type Profile struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Link string `json:"link"`
}

func (p *Profile) SetIDLink(url string) {
	p.Link = strings.ReplaceAll(url, "?viewMode=1", "")
	p.ID = strings.ReplaceAll(p.Link, "https://www.upwork.com/freelancers/~", "")
}
