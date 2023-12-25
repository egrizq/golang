package main

import (
	"fmt"
	"time"
)

func main() {
	// Get the current date and time in the default format.
	currentTime := time.Now()

	// Print the current date and time.
	fmt.Println("Current Date and Time:", currentTime)

	// Extract individual components (year, month, day, etc.).
	year, month, day := currentTime.Date()
	hour, minute, second := currentTime.Clock()

	// Print individual components.
	fmt.Printf("Year: %d, Month: %d, Day: %d\n", year, month, day)
	fmt.Printf("Hour: %d, Minute: %d, Second: %d\n", hour, minute, second)

	// Format the date and time as a string in a specific layout.
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println("Formatted Date and Time:", formattedTime)

	formattedDate := currentTime.Format("02 Jan")
	fmt.Println("Formatted Date:", formattedDate)
}
