package main

import (
	"fmt"
	"time"
)

// Logger simulates a slow log writer
func Logger(logChannel chan string) {
	for msg := range logChannel {
		time.Sleep(500 * time.Millisecond) // Simulate slow I/O
		fmt.Println("Logged:", msg)
	}
}

func main() {
	logChannel := make(chan string, 3) // Buffered channel with capacity 3

	// Start the logger goroutine
	go Logger(logChannel)

	// Simulate rapid log generation
	for i := 1; i <= 5; i++ {
		logMessage := fmt.Sprintf("Log entry #%d", i)
		fmt.Println("Sending:", logMessage)
		logChannel <- logMessage // This won't block until buffer is full
	}

	close(logChannel) // Always close channels when done
}
