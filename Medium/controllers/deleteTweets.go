package controllers

import (
	"job/database"
	"net/http"
	"strconv"
)

// todo delete tweets
func Delete(w http.ResponseWriter, r *http.Request) {
	// parsing form data in the request body
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}

	// todo query the selected id
	selecting, err := database.DB.Query("DELETE FROM tweets WHERE id = ?", id)
	if err != nil {
		panic(err)
	}

	// closed the query
	defer selecting.Close()

	// redirecting the page after deleting the data
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
