package main

import (
	"GORM/database"
	"GORM/model"
	"log"
)

func main() {
	database.Connect()

	// addUser()
	// addData()
	// UpdateData()
	delete()
}

func addUser() {
	// added new user with data

	user := model.User{
		User: "rizq ramadhan",
		News: []model.News{
			{Title: "first News", Main: "Creating first News in the database"},
		},
	}

	if result := database.DB.Create(&user); result.Error != nil {
		panic(result.Error)
	}

	log.Println("new user and data has been added!")
}

func addData() {
	// add data by user id

	id := 1
	var user model.User
	database.DB.Find(&user, id)

	news := model.News{
		Title:  "Second News",
		Main:   "This second data",
		UserID: id,
	}

	user.News = append(user.News, news)

	if result := database.DB.Save(&user.News); result.Error != nil {
		panic(result.Error)
	}

	log.Println("New data has been added!")
}

func UpdateData() {
	// update data

	var data model.News
	id := 1

	database.DB.First(&data, id)
	data.Title = "first update"
	data.Main = "this is the first update"

	if result := database.DB.Save(&data); result.Error != nil {
		panic(result.Error)
	}

	log.Println("success update data!")
}

func delete() {
	// delete data

	var data model.News
	id := 2

	if result := database.DB.Delete(&data, id); result.Error != nil {
		panic(result.Error)
	}

	log.Println("data has been deleted!")
}
