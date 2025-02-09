package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press Return to abort.")
	select {
	case <-time.After(10 * time.Second):
	case <-abort:
		fmt.Println("launch Aborted")
		return

	}
	launch()
}

func launch() {
	fmt.Println("Lift Off")
}
