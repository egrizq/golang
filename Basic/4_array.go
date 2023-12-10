package main

import "fmt"

func main() {
	// array_name := [length]datatype{values} // here length is defined
	list := [5]int{1, 2, 3, 4, 5} // cara pertama deklarasi array
	fmt.Println("Menampilkan array type int:", list)

	list_slice := list[0:3]
	fmt.Println("Slicing array:", list_slice)

	//array_name := [...]datatype{values} // here length is inferred
	list2 := [...]string{"muhammad", "rizq", "ramadhan"}
	fmt.Println("Menampilkan array list2:", list2)

	take := list2[0] //take array
	fmt.Println("Menampilkan list2[0] selection:", take)

	// menghitung array
	fmt.Printf("Panjang array adalah %v", len(list2))

	//iterate array
	for idx, val := range list2 {
		fmt.Printf("iterate list2: %v %v \n", idx, val)
	}
}


