package main

import (
	"fmt"

	"github.com/sony-nurdianto/GoRust/ch12/display"
	"github.com/sony-nurdianto/GoRust/ch12/sexpr"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func main() {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},

		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	display.Display("strangelove", strangelove)

	var i interface{} = 3

	display.Display("i", i)
	display.Display("&i", &i)

	data, err := sexpr.Marshal(strangelove)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Marshal Output:\n", string(data))
	}
}
