package controllers

import (
	"job/database"
	"net/http"
	"time"
)

// todo update data from selected id
func TweetsUpdate(w http.ResponseWriter, r *http.Request) {
	// parsing form data in the request body
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// called time
	currentTime := time.Now()

	// accessing data from form field named
	id_form := r.FormValue("id")
	title_form := r.FormValue("title")
	main_form := r.FormValue("main")
	date_form := currentTime.Format("02 Jan")

	// todo query to update the data
	query := "UPDATE tweets SET title = ?, main = ?, date = ? WHERE id = ?"
	inserting, err := database.DB.Query(query, title_form, main_form, date_form, id_form)
	if err != nil {
		panic(err)
	}

	// closed the query
	defer inserting.Close()

	// todo redirecting the page of HTML after success query
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
