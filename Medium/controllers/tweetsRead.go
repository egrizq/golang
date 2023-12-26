package controllers

import (
	"job/database"
	"job/entities"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

// TODO MENAMPILKAN ISI ARTICLE
func Read(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}

	log.Println(id)
	selecting, err := database.DB.Query("SELECT id, title, main, date FROM tweets WHERE id = ?", id)
	if err != nil {
		panic(err)
	}

	defer selecting.Close()

	var itemSelect []entities.AllTweets

	for selecting.Next() {
		var item entities.AllTweets

		err := selecting.Scan(&item.Id, &item.Title, &item.Main, &item.Date)
		if err != nil {
			panic(err)
		}

		itemSelect = append(itemSelect, item)
	}

	// ganti ke update nanti
	temp, err := template.ParseFiles("views/readTweets.html")
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"Rows": itemSelect,
	}

	temp.Execute(w, data)
}
