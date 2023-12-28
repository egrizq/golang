package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// create a global variabel named DB
var DB *sql.DB

// todo connect go to mysql
func Connect() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3307)/golang")
	if err != nil {
		panic(err)
	}

	DB = db
}
