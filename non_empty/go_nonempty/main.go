package main

import "fmt"

func nonEmpty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}

func nonEmpty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonEmpty(data))

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(arr[1:])
	fmt.Println(arr[2:])
	fmt.Println(arr)
}
