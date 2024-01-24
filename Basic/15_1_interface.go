package main

import "fmt"

type book interface {
	addbook() (string, int)
	databook() (string, int)
}

type bookName struct {
	title string
	qty   int
	place string
	price int
}

func (b bookName) addbook() (string, int) {
	return b.title, b.qty
}

func (p bookName) databook() (string, int) {
	return p.place, p.price
}

func main() {
	var bookCreate book

	bookCreate = bookName{title: "first one", qty: 12, place: "bekasi", price: 12000}

	title, qty := bookCreate.addbook()
	fmt.Printf("book %v have %v stock\n", title, qty)

	place, price := bookCreate.databook()
	fmt.Printf("book %v is %v in %v", title, price, place)
}
