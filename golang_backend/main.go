package main

import (
	"GOLANG_BACKEND/config"
	"GOLANG_BACKEND/controllers"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	http.HandleFunc("/", controllers.Login)
	http.HandleFunc("/auth", controllers.Auth)

	http.HandleFunc("/data", controllers.GetAll)

	http.HandleFunc("/register", controllers.Register)
	http.HandleFunc("/registerSubmit", controllers.RegisterForm)

	http.HandleFunc("/delete", controllers.Delete)

	log.Println("Server run in http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
