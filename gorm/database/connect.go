package database

import (
	"GORM/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/golang"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{}, &model.News{})

	DB = db
}
