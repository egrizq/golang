package main

import "fmt"

func main() {
	type data struct {
		name string
		hp   int
	}

	userdata := data{
		name: "rizq",
		hp:   101010,
	}

	fmt.Println(userdata)
}
