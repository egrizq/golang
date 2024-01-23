package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	// Create a channel to communicate the result back
	resultCh := make(chan int)

	// Use a WaitGroup to wait all goroutines finish
	var wg sync.WaitGroup

	// Lunch 2 goroutines to calculate sum of array concurrently
	wg.Add(3)

	// The 'go' before func are to recognize it is a goroutines
	go calculateArray([]int{12, 5, 3, 2}, resultCh, &wg)
	go calculateArray([]int{10, 2, 4, 1}, resultCh, &wg)
	go calculateArray([]int{22, 1, 4, 1}, resultCh, &wg)

	// Use a goroutines to close the result channel once all calculating
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Receive and sum the results from the channel
	totalSum := 0
	for result := range resultCh {
		totalSum += result
	}

	duration := time.Since(start)
	fmt.Println("total sum:", totalSum, "& time:", duration)
}

func calculateArray(numbers []int, resultCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the WaitGroup counter when the goroutines complete

	sum := 0
	for _, num := range numbers {
		sum += num
	}

	resultCh <- sum // Send the result to the channel
}
