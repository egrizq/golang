package controllers

import (
	"job/database"
	"job/entities"
	"log"
	"net/http"
	"time"
)

func Tweets(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	// time
	currentTime := time.Now()

	title_form := r.FormValue("title")
	main_form := r.FormValue("main")
	date_form := currentTime.Format("02 Jan")

	tweet := entities.AllTweets{
		Title: title_form,
		Main:  main_form,
		Date:  date_form,
	}

	query := "INSERT INTO tweets(title, main, date) VALUES (?, ?, ?)"
	inserting, err := database.DB.Query(query, tweet.Title, tweet.Main, tweet.Date)
	if err != nil {
		panic(err)
	}

	defer inserting.Close()

	log.Println("Inserting:", tweet.Title, tweet.Main, tweet.Date)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
