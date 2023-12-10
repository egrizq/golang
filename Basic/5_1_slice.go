package main

import "fmt"

func main() {

	slice := []int{}

	for n := 0; n < 10; n++ {
		slice = append(slice, n)
	}

	fmt.Println("slice = ", slice)

	read := []int{1, 1, 12, 2, 4, 4, 2, 1}

	for n := 0; n < len(read); n++ {
		if read[n] == 12 {
			fmt.Println("in slice number:", n)
		}
	}
}
