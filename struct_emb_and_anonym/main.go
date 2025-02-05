package main

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Circle struct {
	Point  `json:"point"`
	Radius int `json:"radius"`
}

type Wheel struct {
	Circle `json:"circle"`
	Spokes int `json:"spokes"`
}

func main() {
	wheel := Wheel{
		Circle: Circle{
			Point: Point{
				X: 10,
				Y: 5,
			},
			Radius: 15,
		},
		Spokes: 10,
	}

	json, err := json.MarshalIndent(&wheel, "", "   ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(json))
}
