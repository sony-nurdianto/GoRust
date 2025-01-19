package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()

	args := strings.Join(flag.Args(), *sep)
	if !*n {
		fmt.Print(args)
	} else {
		fmt.Println(args)
	}

}
