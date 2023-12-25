package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3307)/golang")
	if err != nil {
		panic(err)
	}

	DB = db

	log.Println("Database Connected")
}
