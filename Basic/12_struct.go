package main

import "fmt"

func main() {
	type Person struct {
		name string
		age int
		Education string
		salary int
	}

	// variable declare for struct
	var person1 Person
	var person2 Person

	// initialization data
	person1.name = "rizq"
	person1.age = 23
	person1.Education = "software eng"
	person1.salary = 8000000

	person2.name = "james"
	person2.age = 25
	person2.Education = "teacher"
	person2.salary = 3500000

	// calling struct data
	fmt.Printf("nama pertama %v\n", person1.name)
	fmt.Printf("has salary: %v\n", person1.salary) 

	fmt.Printf("nama kedua %v\n", person2.name)
	fmt.Printf("has age of %v\n", person2.age)


	// insert func into struct
	select1 := person1.Education
	selection := category(select1)
	fmt.Printf("rizq are %v\n", selection)
}	

func category(education string) string {
	if education == "string" {
		return ("you get 500")
	} else {
		return ("you get 250")
	}
}