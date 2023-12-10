package main

import "fmt"

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := array[1:4]

	fmt.Println("Array:", array)
	fmt.Println("Slice from array[1:4]:", slice)

	fmt.Println("seleksi slice[0]:", slice[0])

	slice[0] = 20
	fmt.Println("Change value slice[0]:", slice[0])
	fmt.Println("Slice:", slice)
	fmt.Println("Array:", array)

	slice = append(slice, 6)
	fmt.Println("Menambahkan slice[3]: ", slice[3])
	fmt.Println("Slice:", slice)
	fmt.Println("Array:", array)

	fmt.Println("Saat merubah data array melalui slicing maka array dapat berubah juga ")

}
