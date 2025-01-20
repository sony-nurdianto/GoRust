package main

import (
	"fmt"
	"github.com/sony-nurdianto/GoRust/tempconv/go_tempconv/conv"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	for i, v := range args {
		var indexValue = i + 1
		var value float64
		if indexValue < len(args) {
			val, err := strconv.ParseFloat(args[indexValue], 64)
			if err != nil {
				fmt.Printf("Error parse value: %v", err)
				continue
			}
			value = val
		}

		switch v {
		case "-ctf":
			fmt.Printf("Celcius To Farenheit %.2f°F\n", conv.CelciusToFarenheit(value))
		case "-ctk":
			fmt.Printf("Celcius To Kelvin %.2f°K\n", conv.CelciusToKelvin(value))
		case "-ftc":
			fmt.Printf("Farenheit To Celcius %.2f°C\n", conv.FarenheitToCelcius(value))
		case "-ftk":
			fmt.Printf("Farenheit To Kelvin %.2f°K\n", conv.FarenheitToKelvin(value))
		case "-ktc":
			fmt.Printf("Kelvin To Celcius %.2f°C\n", conv.KelvinToCelcius(value))
		case "-ktf":
			fmt.Printf("Kelvin To Fahrenheit %.2f°C\n", conv.KelvinToFarenheit(value))
		default:
			fmt.Println("Unreachable")
		}

		i++
	}
}
