package database

import (
	"gin_gorm/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3307)/golang"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.Data{})

	DB = db
}
