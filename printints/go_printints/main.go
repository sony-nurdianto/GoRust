package main

import (
	"bytes"
	"fmt"
)

func intsToString(value []int) string {
	var buf bytes.Buffer

	buf.WriteByte('[')

	for i, v := range value {
		if i > 0 {
			buf.WriteString(", ")
		}

		fmt.Fprintf(&buf, "%d", v)
	}

	buf.WriteByte(']')
	return buf.String()
}

func main() {
	r := intsToString([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(r)
}
