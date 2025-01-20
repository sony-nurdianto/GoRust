package main

import "fmt"

func main() {
	var x interface{} = 0

	xint, ok := x.(int)
	if !ok {
		fmt.Println("not int")
	} else {
		fmt.Println(xint)
	}
}
