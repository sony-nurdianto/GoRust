package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	counts := make(map[string]int)
	args := os.Args[1:]
	fmt.Println("Find Duplicate Word In " + strings.Join(args, " "))

	for _, filename := range args {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		for _, line := range strings.Split(string(data), "\n") {
			for _, word := range strings.Split(line, " ") {
				re := regexp.MustCompile("[^a-zA-Z0-9]")
				counts[re.ReplaceAllString(strings.TrimSpace(word), "")]++
			}
		}

		for k, v := range counts {
			fmt.Printf("%s: %d\n", k, v)
		}
	}
}
