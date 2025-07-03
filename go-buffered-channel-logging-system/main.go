package main

import (
	"fmt"
	"runtime"
	"time"
)

func Logger(logChannel chan string) {
	for msg := range logChannel {
		time.Sleep(500 * time.Millisecond) // Simulate slow I/O
		fmt.Println("Logged:", msg)
	}
}

func main() {
	runtime.GOMAXPROCS(1) // Force single OS thread to better observe goroutine scheduling

	logChannel := make(chan string, 3) // Buffered channel with capacity 3
	go Logger(logChannel)

	for i := 1; i <= 5; i++ {
		logMessage := fmt.Sprintf("Log entry #%d", i)

		fmt.Printf(">>> Attempting to send: %s\n", logMessage)
		logChannel <- logMessage // Blocks after buffer is full
		fmt.Printf("✓✓✓ Sent: %s\n", logMessage)
	}

	close(logChannel)
}
