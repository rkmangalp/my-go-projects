package main

import (
	"fmt"
	"time"
)

// Worker function
func worker(id int, ch chan string) {
	time.Sleep(2 * time.Second) // Simulating work
	ch <- fmt.Sprintf("Worker %d finished", id)
}

func main() {
	ch := make(chan string) // Creating a channel

	// Launching multiple goroutines
	for i := 1; i <= 3; i++ {
		go worker(i, ch)
	}

	// Receiving results from goroutines
	for i := 1; i <= 3; i++ {
		fmt.Println(<-ch) // Blocking call until a worker sends data
	}
}
