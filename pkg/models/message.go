package models

type MessageUser struct {
	UserName     string  `json:"username"`
	Password     string  `json:"password"`
	SecretAwnser string  `json:"secret_awnser"`
	ProfileData  Profile `json:"extra_data"`
	Retries      int64   `json:"retries_count"`
}
