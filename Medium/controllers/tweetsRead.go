package controllers

import (
	"job/database"
	"job/entities"
	"net/http"
	"strconv"
	"text/template"
)

// todo showing the main article are selected from main page
func Read(w http.ResponseWriter, r *http.Request) {
	// todo GET the 'id' from HTML
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}

	// todo query the id to show id, title, main, and date that selected
	selecting, err := database.DB.Query("SELECT id, title, main, date FROM tweets WHERE id = ?", id)
	if err != nil {
		panic(err)
	}

	// closed the query
	defer selecting.Close()

	// create a slice to store the query/'selecting' data
	var itemSelect []entities.AllTweets

	// .Next() are iterate the query
	for selecting.Next() {
		// create a variable to save the data in query
		var item entities.AllTweets

		// .Scan() used to scan the values from the row of query into a variable/ 'item'
		err := selecting.Scan(&item.Id, &item.Title, &item.Main, &item.Date)
		if err != nil {
			panic(err)
		}

		// append 'item' into the slice of the 'itemSelect'
		itemSelect = append(itemSelect, item)
	}

	// todo template the html
	temp, err := template.ParseFiles("views/readTweets.html")
	if err != nil {
		panic(err)
	}

	// store the 'itemSelect' to a map as a values, with Rows as a key
	data := map[string]interface{}{
		"Rows": itemSelect,
	}

	// todo excute the HTML with 'data'
	temp.Execute(w, data)
}
