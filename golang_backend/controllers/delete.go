package controllers

import (
	"GOLANG_BACKEND/config"
	"net/http"
	"strconv"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		http.Error(w, "Invalid or missing ID parameter", http.StatusBadRequest)
	}

	deleting, err := config.DB.Query("DELETE FROM mydata WHERE id = ?", id)
	if err != nil {
		panic(err)
	}

	defer deleting.Close()
	http.Redirect(w, r, "/data", http.StatusSeeOther)

}
