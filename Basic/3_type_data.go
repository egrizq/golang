package main

import "fmt"

var nama string = "rizq"

func main() {
	var (
		integer  = 12
		floating = 12.90
		boolean  = true
		String  = "Hello"
	)

	fmt.Printf("type: %T\n", integer)
	fmt.Printf("type: %T\n", floating)
	fmt.Printf("type: %T\n", boolean)
	fmt.Printf("type: %T\n", String)
	fmt.Printf("namaku: %v", nama)
}
