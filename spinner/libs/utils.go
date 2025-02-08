package libs

import (
	"fmt"
	"time"
)

func Spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c\n", r)
			time.Sleep(delay)
		}
	}
}

func Fib(x int) int {
	if x < 2 {
		return x
	}
	return Fib(x-1) + Fib(x-2)
}
