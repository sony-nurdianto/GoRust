package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

type Kv struct {
	Key   string
	Value int
}

func main() {
	args := os.Args[1:]

	for _, filename := range args {

		fmt.Println("Find Duplicate Word In " + filename)
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		counts := make(map[string]int)

		for _, line := range strings.Split(string(data), "\n") {
			for _, word := range strings.Split(line, " ") {
				re := regexp.MustCompile("[^a-zA-Z0-9]")
				counts[re.ReplaceAllString(strings.TrimSpace(word), "")]++
			}
		}

		var sortedCount []Kv

		for k, v := range counts {
			sortedCount = append(sortedCount, Kv{Key: k, Value: v})
		}

		sort.Slice(sortedCount, func(i, j int) bool {
			return sortedCount[i].Value > sortedCount[j].Value
		})

		for _, pair := range sortedCount {
			fmt.Printf("%s: %d\n", pair.Key, pair.Value)
		}
	}
}
