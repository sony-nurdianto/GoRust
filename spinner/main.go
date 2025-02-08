package main

import (
	"fmt"
	"time"

	"github.com/sony-nurdianto/GoRust/spinner/libs"
)

func main() {
	go libs.Spinner(100 * time.Millisecond)
	const n = 45
	fibN := libs.Fib(n)
	fmt.Printf("\rFibbonaci(%d) = %d\n", n, fibN)
}
