package controllers

import (
	"GOLANG_BACKEND/config"
	"log"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	username_insert := r.FormValue("username")
	password_insert := r.FormValue("password")

	log.Println("username:", username_insert, "and password:", password_insert)

	selecting, err := config.DB.Query("SELECT username, password FROM mydata WHERE username = ? and password = ?", username_insert, password_insert)
	if err != nil {
		panic(err)
	}

	for selecting.Next() {
		var username string
		var password string

		err := selecting.Scan(&username, &password)
		if err != nil {
			panic(err)
		}

		if (username == '') && (password == '') {
			
		}

		// if id ==  {
		// 	log.Println("Login Success")
		// 	http.Redirect(w, r, "/data", http.StatusSeeOther)
		// } else {
		// 	log.Println("Login Failed")
		// 	w.Header().Set("Content-Type", "text/html")
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	fmt.Fprint(w, `<html><body><script>alert("Login failed!");</script></body></html>`)
		// }
	}
}
