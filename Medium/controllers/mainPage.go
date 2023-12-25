package controllers

import (
	"job/database"
	"job/entities"
	"log"
	"net/http"
	"text/template"
)

func ForYou(w http.ResponseWriter, r *http.Request) {
	rows, err := database.DB.Query("SELECT * FROM tweets")
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	var all []entities.AllTweets

	for rows.Next() {
		var row entities.AllTweets

		err := rows.Scan(&row.Id, &row.Title, &row.Main, &row.Date)
		if err != nil {
			panic(err)
		}

		all = append(all, row)
	}

	temp, err := template.ParseFiles("views/index.html")
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"Rows": all,
	}

	log.Println("Main Page")
	temp.Execute(w, data)
}
