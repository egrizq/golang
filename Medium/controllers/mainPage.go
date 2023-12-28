package controllers

import (
	"job/database"
	"job/entities"
	"net/http"
	"text/template"
)

// todo showing the data to main page
func ForYou(w http.ResponseWriter, r *http.Request) {
	// todo query all the data from database
	rows, err := database.DB.Query("SELECT * FROM tweets")
	if err != nil {
		panic(err)
	}

	// closed the query
	defer rows.Close()

	// store the data into slice
	var allItem []entities.AllTweets

	// .Next() are iterate the query
	for rows.Next() {
		// create a variable to save the data
		var row entities.AllTweets

		// .Scan() used to scan the values from the row of query into a variable/ 'row'
		err := rows.Scan(&row.Id, &row.Title, &row.Main, &row.Date)
		if err != nil {
			panic(err)
		}

		// append 'item' into the slice of the 'itemSelect'
		allItem = append(allItem, row)
	}

	// todo templating the data into main page == index.html
	temp, err := template.ParseFiles("views/index.html")
	if err != nil {
		panic(err)
	}

	// store the 'allItem' to a map as a values, with Rows as a key
	data := map[string]interface{}{
		"Rows": allItem,
	}

	// todo excute data to html
	temp.Execute(w, data)
}
