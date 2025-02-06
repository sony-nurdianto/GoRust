package main

import "fmt"

func sum(vals ...int) int {
	total := 0

	for _, val := range vals {
		total += val
	}

	return total
}

func main() {
	total := sum(2, 2, 2, 2, 2)

	fmt.Println(total)
}
