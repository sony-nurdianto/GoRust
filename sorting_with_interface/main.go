package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (a StringSlice) Len() int           { return len(a) }
func (a StringSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a StringSlice) Less(i, j int) bool { return a[i] < a[j] }

func main() {
	data := StringSlice{
		"Durrian",
		"Apple",
		"Cerry",
		"Berry",
	}

	sort.Sort(data)
	fmt.Println(data)
}
