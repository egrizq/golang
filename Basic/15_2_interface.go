package main

import "fmt"

type shoes interface {
	shoesName() []string
}

type name struct {
	name []string
}

func (n name) shoesName() []string {
	return n.name
}

func main() {
	var listShoes shoes

	listShoes = name{[]string{"adidas", "nike", "new balance"}}
	for _, one := range listShoes.shoesName() {
		fmt.Println("shoes:", one)
	}
}
