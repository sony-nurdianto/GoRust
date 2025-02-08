package main

import "fmt"

func main() {
	natruals := make(chan int)
	square := make(chan int)

	go func() {
		for x := 0; ; x++ {
			natruals <- x
		}
	}()

	go func() {
		for {
			x := <-natruals
			square <- x * x
		}
	}()

	for {
		fmt.Println(<-square)
	}
}
