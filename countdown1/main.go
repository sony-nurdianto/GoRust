package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Commencing Countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println("Countdown")
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("Lift Off")
}
