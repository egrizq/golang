package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	say := "hello"
	to := "to me"
	// %v calling value
	// %#v calling the value
	fmt.Printf("say %v %#v", say, to) // calling variabel with Printf
}
