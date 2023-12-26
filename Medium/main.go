package main

import (
	"job/controllers"
	"job/database"
	"log"
	"net/http"
)

func main() {
	database.Connect()

	http.HandleFunc("/", controllers.ForYou) // menampilkan semua arcticle yang ada

	http.HandleFunc("/tweets", controllers.TweetsForm)     // input article
	http.HandleFunc("/tweets/publish", controllers.Tweets) // publish article via form

	http.HandleFunc("/tweets/read", controllers.Read)                // membaca 1 article dalam For You
	http.HandleFunc("/tweets/update", controllers.Update)            // menampilkan 1 article dalam bentuk input
	http.HandleFunc("/tweets/update/form", controllers.TweetsUpdate) // update article via form

	http.HandleFunc("/tweets/delete", controllers.Delete) // delete tweets

	log.Println("Server Run in http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
