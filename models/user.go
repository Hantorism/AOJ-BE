package models

type User struct {
	Email         string `json:"email"`
	Nickname      string `json:"nickname"`
	StudentNumber string `json:"studentNumber"`
	Password      string `json:"password"`
}
