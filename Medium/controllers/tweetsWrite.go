package controllers

import (
	"log"
	"net/http"
	"text/template"
)

func TweetsWrite(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/tweets.html")
	if err != nil {
		panic(err)
	}

	log.Println("Input tweets...")
	temp.Execute(w, nil)
}
