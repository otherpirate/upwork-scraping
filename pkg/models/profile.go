package models

import "strings"

type Address struct {
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	City       string `json:"city"`
	State      string `json:"state"`
	PostalCode string `json:"postal_code"`
	Country    string `json:"country"`
}

type Employment struct {
	Status              string `json:"status"`
	Type                string `json:"type"`
	JobTitle            string `json:"job_title"`
	PlatformUserID      string `json:"platform_user_id"`
	HireDatetime        string `json:"hire_datetime"`
	TerminationDatetime string `json:"termination_datetime"`
	TerminationReason   string `json:"termination_reason"`
}

type Profile struct {
	ID            string     `json:"id"`
	Account       string     `json:"account"`
	Employer      string     `json:"employer"`
	CreatedAt     string     `json:"created_at"`
	UpdatedAt     string     `json:"updated_at"`
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	FullName      string     `json:"full_name"`
	Email         string     `json:"email"`
	PhoneNumber   string     `json:"phone_number"`
	BirthDate     string     `json:"birth_date"`
	PictureURL    string     `json:"picture_url"`
	Address       Address    `json:"address"`
	Employment    Employment `json:"employment"`
	SSN           string     `json:"ssn"`
	MaritalStatus string     `json:"marital_status"`
	Gender        string     `json:"gender"`
}

func (p *Profile) SetNames(name string) {
	p.FullName = name
	names := strings.Split(name, " ")
	p.FirstName = names[0]
	if len(names) > 0 {
		p.LastName = names[1]
	}
}
