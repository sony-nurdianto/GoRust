// Dup2 prints the count and text of lines that appear
// more than once in input. It reads from stdin or from a list
// of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	if err := input.Err(); err != nil {
		fmt.Println("Error do scan:", err)
		return
	}

	for input.Scan() {
		line := strings.Split(input.Text(), " ")

		for _, word := range line {
			re := regexp.MustCompile("[^a-zA-Z0-9]")
			counts[re.ReplaceAllString(strings.TrimSpace(word), "")]++
		}

	}

	for key, value := range counts {
		fmt.Printf("%s: %d\n", key, value)
	}
}

func main() {
	files := os.Args[1:]
	counts := make(map[string]int)

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		f, err := os.Open(files[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		defer f.Close()
		countLines(f, counts)
	}
}
