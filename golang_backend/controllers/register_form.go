package controllers

import (
	"GOLANG_BACKEND/config"
	"log"
	"net/http"
	"strconv"
)

type UserData struct {
	Username  string
	Email     string
	Handphone string
	Password  string
}

func RegisterForm(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	username_form := r.FormValue("username")
	email_form := r.FormValue("email")
	handphone_form := r.FormValue("handphone")
	password_form := r.FormValue("password")

	handphone, err := strconv.Atoi(handphone_form)
	if err != nil {
		panic(err)
	}

	userData := UserData{
		Username:  username_form,
		Email:     email_form,
		Handphone: handphone_form,
		Password:  password_form,
	}

	dataReady := userData.Username != "" || userData.Email != "" || handphone != 0 || userData.Password != ""

	if dataReady {
		log.Println("Data is ready")
	} else {
		log.Println("Data are empty")
	}

	query := "INSERT INTO mydata(username, email, handphone, password) VALUES (?, ?, ?, ?)"
	inserting, err := config.DB.Query(query, userData.Username, userData.Email, handphone, userData.Password)
	if err != nil {
		panic(err)
	}

	defer inserting.Close()

	log.Println("Success inserting into mydata")
	http.Redirect(w, r, "/database", http.StatusSeeOther)
}
