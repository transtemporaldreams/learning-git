package main

import (
	"fmt"
	"time"
)

func main() {
	// print the current time
	now := time.Now().UTC()

	fmt.Printf("the time is: %s\n", now)
	fmt.Printf("formatted: %s\n", now.Format(time.RFC3339))
}