package main

import "fmt"

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

type Point struct {
	X, Y int
}

func main() {
	table := make(map[string]int)

	table["one"] = 1
	table["two"] = 2
	table["three"] = 3

	s := k([]string{"one", "two", "three"})
	fmt.Println(s)

	point := Point{1, 2}
	fmt.Println(point)
}
