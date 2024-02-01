package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type User struct {
	ID       uint   `form:"id" gorm:"primarykey"`
	Username string `form:"username"`
	Password string `form:"password"`
}

func Connect() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/golang"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})

	DB = db
}
