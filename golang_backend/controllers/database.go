package controllers

import (
	"GOLANG_BACKEND/config"
	"html/template"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(`SELECT * FROM mydata`)
	if err != nil {
		panic(err)
	}

	defer rows.Close()

	type TableRow struct {
		Id        int
		Username  string
		Email     string
		Handphone int
		Password  string
	}

	var tableRows []TableRow

	for rows.Next() {
		var row TableRow

		err := rows.Scan(&row.Id, &row.Username, &row.Email, &row.Handphone, &row.Password)
		if err != nil {
			panic(err)
		}

		tableRows = append(tableRows, row)

	}

	temp, err := template.ParseFiles("views/show/show.html")
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"Rows": tableRows,
	}

	temp.Execute(w, data)
}
