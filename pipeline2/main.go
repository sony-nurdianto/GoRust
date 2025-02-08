package main

import "fmt"

func main() {
	natruals := make(chan int)
	square := make(chan int)

	go func() {
		for x := 0; x < 100; x++ {
			natruals <- x
		}
		close(natruals)
	}()

	go func() {
		for {
			x, ok := <-natruals
			if !ok {
				break
			}
			square <- x * x
		}
		close(square)
	}()

	for x := range square {
		fmt.Println(x)
	}
}
