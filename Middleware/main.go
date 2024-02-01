package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// todo project ini menggunakan postman

// Definisi dari middleware sendiri versi penulis,
// sebuah blok kode yang dipanggil sebelum ataupun sesudah http request di proses.

// *note: handlerFunc adalah type, sedangkan handleFunc adalah function.
// *handlerFunc digunakan untuk wrapping function menjadi type yang satisfies interface http.Handler.

func Loggin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if the method is GET then check the id
		if r.Method == "GET" {
			// get the id from route
			id, err := strconv.Atoi(r.URL.Query().Get("id"))
			if err != nil {
				panic(err)
			}
			log.Println(id)

			// if the id not 10 then redirect
			if id != 10 {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}

			// serve the endpoint if ok
			next.ServeHTTP(w, r)

			// if the method is POST then proceed with simple login
		} else if r.Method == "POST" {
			// define the struct for body in json type
			var Data struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}

			// parse the json body
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}

			// convert the json body into struct
			err = json.Unmarshal(body, &Data)
			if err != nil {
				log.Println("cannot parse body")
			}

			log.Println(Data.Username, Data.Password)

			// check the login and serve if correct
			if Data.Username == "rizq" && Data.Password == "rizq" {
				log.Println("success")
				next.ServeHTTP(w, r)
			} else {
				log.Println("wrong data")
				http.Redirect(w, r, "/", http.StatusSeeOther)
			}
		}

	})
}

func direct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "direct")
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func woo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "woo")
}

func main() {
	http.HandleFunc("/", direct)
	http.HandleFunc("/foo", Loggin(foo))
	http.HandleFunc("/woo", Loggin(woo))

	http.ListenAndServe(":8000", nil)
}
