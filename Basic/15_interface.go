package main

import (
	"fmt"
)

// todo define the interface
type calc interface {
	addition() float64 // method
	multiple() float64 // method
}

// struct as type in method
type number struct {
	width, height float64
}

func (n number) addition() float64 {
	return n.width + 10
}

func (n number) multiple() float64 {
	return n.width * n.height
}

func main() {
	var hitung calc

	hitung = number{12, 10}
	fmt.Println("addition:", hitung.addition())
	fmt.Println("multiple:", hitung.multiple())

	hitung = number{22, 30}
	fmt.Println("addition:", hitung.addition())
	fmt.Println("multiple:", hitung.multiple())
}
