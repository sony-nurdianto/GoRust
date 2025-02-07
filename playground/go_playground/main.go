package main

import (
	"fmt"
)

type Voice interface {
	Speak() string
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "Hallo, My Name is " + p.Name
}

type Robot struct {
	Name string
}

func (r Robot) Speak() string {
	return "Beep Bop, My Name Is " + r.Name
}

func main() {
	var voice Voice

	voice = Person{Name: "Sony"}
	fmt.Println(voice.Speak())

	voice = Robot{Name: "X01"}
	fmt.Println(voice.Speak())
}
