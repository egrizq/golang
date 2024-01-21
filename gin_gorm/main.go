package main

import (
	"gin_gorm/controllers"
	"gin_gorm/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	database.Connect()

	router.LoadHTMLGlob("templates/*")
	router.Static("/public", "./public")

	router.GET("/", controllers.Login)                // views login
	router.POST("/login/form", controllers.LoginForm) // login

	router.GET("/add", controllers.Views)
	router.GET("/dashboard", controllers.Dashboard)
	router.GET("/dashboard/edit", controllers.Edit)
	router.GET("/dashboard/delete", controllers.Delete)
	router.GET("/logout", controllers.ClearSession)
	router.POST("/dashboard/edit/update", controllers.UpdateData)
	router.POST("/create", controllers.Create)

	router.Run(":8000")
}
