package main

import "fmt"

func main() {
	for dev := 1; dev <= 10; dev += 2 {
		//fmt.Println(dev)
	}

	data := "indonesia"
	for n := 0; n < len(data); n++ {
		//fmt.Println(string(data[n]))
	}

	// i access index of each character
	// item access each character
	list := [4]int{1, 2, 3, 4}
	for i, item := range list {
		fmt.Println(i, item)
	}

}
