package main

import "fmt"

func main() {
	var funcs []func()

	// Loop biasa
	for i := 0; i < 5; i++ {
		funcs = append(funcs, func() {
			fmt.Println("Value of i:", i)
		})
	}

	fmt.Println("Executing functions:")
	for _, f := range funcs {
		f()
	}
}
