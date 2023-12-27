package controllers

import (
	"job/database"
	"log"
	"net/http"
	"time"
)

// TODO FORM HTML UPDATE
func TweetsUpdate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// time
	currentTime := time.Now()

	id_form := r.FormValue("id")
	title_form := r.FormValue("title")
	main_form := r.FormValue("main")
	date_form := currentTime.Format("02 Jan")

	query := "UPDATE tweets SET title = ?, main = ?, date = ? WHERE id = ?"
	inserting, err := database.DB.Query(query, title_form, main_form, date_form, id_form)
	if err != nil {
		panic(err)
	}

	defer inserting.Close()

	log.Println("Update:", id_form, title_form, main_form, date_form)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
