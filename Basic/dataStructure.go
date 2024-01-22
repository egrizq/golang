package main

import (
	"fmt"
)

func main() {
	fibonacci()
	fizzbuzz()
}

func fibonacci() {
	var loop int = 0
	var firstNumber int = 0
	var secondNumber int = 1

	var handleEven int = 0
	var handleOdd int = 0

	const max int = 4_000_000

	for loop <= max {
		var fibo int = firstNumber
		firstNumber = secondNumber
		secondNumber = firstNumber + fibo

		if secondNumber%2 == 0 {
			handleEven += secondNumber
		} else {
			handleOdd += secondNumber
		}

		loop = secondNumber + firstNumber
	}

	fmt.Println("total even:", handleEven)
	fmt.Println("total odd:", handleOdd)
}

func fizzbuzz() {
	var temp int = 0

	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			temp += i
		} else if i%5 == 0 {
			temp += i
		}
	}

	fmt.Println("total number:", temp)
}
