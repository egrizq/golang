package main

import (
	"session/controllers"
	"session/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")
	// middleware to all the routes
	router.Use(controllers.Logger())

	database.Connect()

	// if "/" doesn't have a session then go to "/login"
	router.GET("/", controllers.Session(), controllers.Main)
	router.GET("/login", controllers.Login)

	// routes POST for form login, if correct go to "/"
	router.POST("/login/form", controllers.LoginPostHandler)

	router.Run(":8000")
}
