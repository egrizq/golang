package main

import "fmt"

func calculate(x int, y int) (int, int) {
	var kali = x * y
	var bagi = x / y

	return kali, bagi
}

func namaAsal(depan string, asal_input string) (nama_depan string, asal string) {
	nama_depan = depan
	asal = asal_input

	return
}

func main() {
	kali, bagi := calculate(10, 5)

	fmt.Println("kali:", kali)
	fmt.Println("bagi:", bagi)

	nama_depan, asal := namaAsal("rizq", "bekasi")

	fmt.Println(nama_depan, "dan", asal)
}
