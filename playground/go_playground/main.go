package main

import "fmt"

func sum(n int) int {
	fmt.Println(n)
	if n > 10 {
		return n
	}

	return sum(n + 1)
}

func main() {
	n := 1
	s := sum(n)
	fmt.Printf("This is Rekursif Result sum(%d) : %d\n", n, s)
}
