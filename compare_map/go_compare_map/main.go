package main

import "fmt"

func equals(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}

	return true
}

func main() {
	xmap := map[string]int{"some": 1, "none": 0}
	ymap := map[string]int{"some": 1, "none": 0}

	for k, v := range xmap {
		fmt.Printf("%s:%d\n", k, v)
	}

	fmt.Println(len(xmap))

	isEqueal := equals(xmap, ymap)

	fmt.Println(isEqueal)
}
