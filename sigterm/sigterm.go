package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	handleInterrupt()

	// Do something pointless until interrupted
	for i := 1; ; i++ {
		fmt.Printf("%v mississippi...\n", i)
		time.Sleep(1 * time.Second)
	}
}

func cleanup() {
	fmt.Println("Bye!")
}

// Call cleanup() when interrupted before cleanly exiting
func handleInterrupt() {
	sigterm := make(chan os.Signal)
	signal.Notify(sigterm, os.Interrupt)

	go func() {
		<-sigterm
		cleanup()
		os.Exit(1)
	}()
}
