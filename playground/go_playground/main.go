package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 600 // canvas size
	radius        = 200      // radius dari dasar kerucut
	heightCone    = 300      // tinggi kerucut
)

func main() {
	// Membuat file SVG output
	file, err := os.Create("cone.svg")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Menulis header SVG
	fmt.Fprintf(file, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: black; fill: transparent; stroke-width: 1' "+
		"width='%d' height='%d'>", width, height)

	// Titik pusat dasar kerucut
	centerX := width / 2
	centerY := height - 50 // Sedikit lebih dari bawah canvas

	// Titik puncak kerucut
	topX := width / 2
	topY := height - heightCone

	// Membuat dasar kerucut (lingkaran)
	for i := 0; i < 360; i++ {
		angle := float64(i) * math.Pi / 180 // Konversi derajat ke radian
		x := float64(centerX) + radius*math.Cos(angle)
		y := float64(centerY) - radius*math.Sin(angle)

		// Menghubungkan titik pada lingkaran dengan puncak kerucut
		fmt.Fprintf(file, "<line x1='%f' y1='%f' x2='%f' y2='%f' />\n", float64(topX), float64(topY), x, y)
	}

	// Menutup tag SVG
	fmt.Fprintf(file, "</svg>")

	fmt.Println("SVG file generated: cone.svg")
}
