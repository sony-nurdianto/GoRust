package main

import (
	"fmt"
	"time"
)

func count(p *int) {
	*p++
}

func main() {
	start := time.Now()

	var p int

	for i := 1; i <= 10; i++ {
		count(&p)
	}

	fmt.Printf("%d\n", p)

	elapsed := time.Since(start)
	fmt.Printf("Elapsed Time: %v\n", elapsed)
}
