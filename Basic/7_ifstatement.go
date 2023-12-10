package main

import "fmt"

func main() {
	umur := 12
	prestasi := true

	if umur >= 12 {
		fmt.Println("Lolos")
	} else if umur <= 12 {
		fmt.Println("Tidak Lolos")
	} else if umur <= 12 && prestasi == true {
		fmt.Println("Lolos jalur prestasi")
	} else {
		fmt.Println("Tidak lolos")
	}
}
