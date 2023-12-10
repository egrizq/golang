package main

import "fmt"

func main() {
	game := true

	for game {
		fmt.Println("True")
		for n := 0; n < 10; n++{
			if n == 10 {
				break
			}
		}
	}
}
