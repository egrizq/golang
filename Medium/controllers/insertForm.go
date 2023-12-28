package controllers

import (
	"job/database"
	"net/http"
	"time"
)

// todo inserting a new tweets
func TweetsPublish(w http.ResponseWriter, r *http.Request) {
	// parsing form data in the request body
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// called time
	currentTime := time.Now()

	// accessing data from form field named
	title_form := r.FormValue("title")
	main_form := r.FormValue("main")
	date_form := currentTime.Format("02 Jan")

	// todo query to insert the data
	query := "INSERT INTO tweets(title, main, date) VALUES (?, ?, ?)"
	inserting, err := database.DB.Query(query, title_form, main_form, date_form)
	if err != nil {
		panic(err)
	}

	// closed the query
	defer inserting.Close()

	// redirecting the page after inserting a new tweet
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
