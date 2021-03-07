package models

type MessageUser struct {
	UserName     string `json:"username"`
	Password     string `json:"password"`
	SecretAwnser string `json:"secret_awnser"`
}
