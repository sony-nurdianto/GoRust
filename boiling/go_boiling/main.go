package main

import (
	"errors"
	"fmt"
)

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(boilingF))

	if err := doSomething("break"); err != nil {
		fmt.Println(err)
	}

}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func doSomething(act string) error {
	switch act {
	case "action":
		fmt.Println("Action !!!!")
	case "cut":
		fmt.Println("CUT !!!!")
	case "break":
		fmt.Println("Good Jobs !!!!")
	default:
		return errors.New("You Peace of a shit !!!!")
	}
	return nil
}
