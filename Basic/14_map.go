package main

import "fmt"

func main() {
	setGame := map[int]string{1:"baseball", 2:"football", 3:"volley"}

	fmt.Println("this is a map:", setGame)
	fmt.Println("indexing setGame 1", setGame[1])

	// looping setGame
	for idx, val := range setGame {
		fmt.Printf("game %v is a %v\n", idx, val)
	}

	// add new value 
	setGame[4] = "tennis"
	fmt.Println("adding value 4", setGame)

	// delete value
	delete(setGame, 1)
	fmt.Println("removing value 1", setGame)

}