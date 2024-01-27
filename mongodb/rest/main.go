package main

import (
	"rest/controllers"
	"rest/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()
	router := gin.Default()

	router.GET("/shows", controllers.AllData)

	router.Run(":8000")
}
