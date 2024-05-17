package models

type User struct {
	Id            int    `json:"id"`
	Email         string `json:"email"`
	Nickname      string `json:"nickname"`
	StudentNumber string `json:"student_number"`
	Password      string `json:"password"`
}
