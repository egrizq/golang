package controllers

import (
	"job/database"
	"net/http"
	"strconv"
)

// TODO DELETE TWEETS
func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}

	selecting, err := database.DB.Query("DELETE FROM tweets WHERE id = ?", id)
	if err != nil {
		panic(err)
	}

	defer selecting.Close()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
