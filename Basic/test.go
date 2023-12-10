package main

import "fmt"

func main() {

	nama := "muhammmad rizq ramadhan"
	const AGE int = 13

	println(nama, " is ", AGE, "years old")

	year := [...]int{1, 2, 3, 4}
	kelas := [...]string{"sd", "smp", "smk", "kuliah"}

	for _, val := range year {
		for _, val1 := range kelas {
			if val == 1 || val1 == "sd" {
				fmt.Printf("index %v dengan value %v \n", val, val1)
			}
		}
	}
	
	sumTest(3,2)
}


func sumTest(first int, second int) {
	list := [...]int{10,3,7,5,3}

	first_return := ""
	second_return := ""

	for i :=0 ; i < len(list); i++ {
		if i == first {
			first_return = "true"
		}

		if i == second {
			second_return = "true"
		}
	}

	fmt.Println(first_return, second_return)
}