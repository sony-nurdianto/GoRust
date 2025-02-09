package main

import "fmt"

func mirroredQuery() string {
	responses := make(chan string, 3)
	go func() { responses <- "asia.gopl.io" }()
	go func() { responses <- "europe.gopl.io" }()
	go func() { responses <- "us.gopl.io" }()
	return <-responses
}

func main() {
	fmt.Println(mirroredQuery())
}
