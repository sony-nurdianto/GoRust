package main

import (
	"fmt"

	"github.com/sony-nurdianto/GoRust/space_math/geometry"
)

func main() {
	q := geometry.Point{X: 1, Y: 2}
	p := geometry.Point{X: 4, Y: 6}

	fmt.Println(geometry.Distance(p, q))
	p.ScaleBy(2.0)
	fmt.Println(p.Distance(q))

	perim := geometry.Path{
		geometry.Point{
			X: 1, Y: 1,
		},
		geometry.Point{
			X: 5, Y: 1,
		},
		geometry.Point{
			X: 5, Y: 4,
		},
		geometry.Point{
			X: 1, Y: 1,
		},
	}

	fmt.Println(perim.Distance())
}
