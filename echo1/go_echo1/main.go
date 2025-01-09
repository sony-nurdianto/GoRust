package main

import (
	"fmt"
	"os"
	"strings"
)

func MyEcho(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func ExampleEcho(args []string) {
	var s, sep string
	for i := 1; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

func main() {
	var args []string = os.Args[1:]
	MyEcho(args)
	ExampleEcho(args)
}
