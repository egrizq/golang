package controllers

import (
	"net/http"
	"text/template"
)

func TweetsWrite(w http.ResponseWriter, r *http.Request) {
	// todo templating 'Home Page' in HTML
	temp, err := template.ParseFiles("views/tweets.html")
	if err != nil {
		panic(err)
	}

	// todo excute the HTML file
	temp.Execute(w, nil)
}
