package main

import (
	"fmt"
	"time"
)

func main() {
	// Get current local time
	currentTime := time.Now()

	// Print in default format
	fmt.Println("Current Time:", currentTime)

	// Print in custom format
	fmt.Println("Formatted:", currentTime.Format("01-02-2006 15:04:05 Monday"))
}
