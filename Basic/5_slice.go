package main

import "fmt"

func main() {
	// slice_name := []datatype{values}
	// [] is unlimited and can be modify
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("slice:", slice)
	fmt.Println("return of length:", len(slice))

	//returns the capacity of the slice (the number of elements the slice can grow or shrink to)
	fmt.Println("capacity:", cap(slice))

	fmt.Println("slicing slice[2:]:", slice[2:])

	slice = append(slice, 6)
	fmt.Println("menambahkan value ke slice dengan value:", slice[5])
	fmt.Println("Slice: ", slice)
}
