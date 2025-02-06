package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sony-nurdianto/GoRust/findlinks3/links"
)

func breadFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist

		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)

	list, err := links.ExtractUrl(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	breadFirst(crawl, os.Args[1:])
}
