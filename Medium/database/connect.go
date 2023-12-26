package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

var DB *sql.DB

// type Tweets struct {
// 	gorm.Model
// 	id    string
// 	title string
// 	main  string
// 	date  string
// }

func Connect() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3307)/golang")
	if err != nil {
		panic(err)
	}

	DB = db

	log.Println("Database Connected")
}

// func Connect() {
// 	var err error

// 	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3307)/golang"), &gorm.Config{})
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = db.AutoMigrate(&Tweets{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
