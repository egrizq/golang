package controllers

import (
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/login/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, nil)
}
