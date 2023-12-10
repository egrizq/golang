package main

import "fmt"

func setPlay(name string) {
	fmt.Println("you're in the game", name)
}

func goBack(name string, prize int) {
	fmt.Printf("you can't play %v and here's the prize %v", name, prize)
}

func play(playing bool, name string) {
	if playing == true {
		setPlay(name)
	} else if playing == false {
		goBack(name, 5000)
	}
}

func main() {
	play(false, "rizq")
}