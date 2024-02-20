package models

type Users struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type StoreUser struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
}
