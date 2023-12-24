package controllers

import (
	"html/template"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/register/register.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
