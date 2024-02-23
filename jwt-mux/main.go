package main

import (
	"jwt-mux/controllers"
	"jwt-mux/database"
	"jwt-mux/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	database.ConnectDB()
	router := mux.NewRouter()

	router.HandleFunc("/login", controllers.Login).Methods("POST")
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/logout", controllers.Logout).Methods("GET")

	data := router.PathPrefix("/get").Subrouter()
	data.HandleFunc("/alldata", controllers.Alldata).Methods("GET")
	data.Use(middleware.JWTmiddle)

	log.Println("Server is running on http://localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
