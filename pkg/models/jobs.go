package models

type Job struct {
	Title           string   `json:"title"`
	Resume          string   `json:"resume"`
	Link            string   `json:"link"`
	Type            string   `json:"type"`
	Level           string   `json:"level"`
	Estimate        string   `json:"estimate"`
	PostedAt        string   `json:"posted_at"`
	Skills          []string `json:"skills"`
	Proposals       string   `json:"proposals"`
	PaymentVerified bool     `json:"payment_verified"`
	Spent           string   `json:"spent"`
	Location        string   `json:"location"`
}
