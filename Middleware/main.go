package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Definisi dari middleware sendiri versi penulis,
// sebuah blok kode yang dipanggil sebelum ataupun sesudah http request di proses.

// *note: handlerFunc adalah type, sedangkan handleFunc adalah function.
// *handlerFunc digunakan untuk wrapping function menjadi type yang satisfies interface http.Handler.

func Loggin(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			id, err := strconv.Atoi(r.URL.Query().Get("id"))
			if err != nil {
				panic(err)
			}
			log.Println(id)
			if id != 10 {
				http.Redirect(w, r, "/", http.StatusSeeOther)
				return
			}

			next.ServeHTTP(w, r)
		} else if r.Method == "POST" {
			next.ServeHTTP(w, r)
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
