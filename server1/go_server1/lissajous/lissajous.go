package lissajous

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"strconv"
)

var palette = []color.Color{
	color.White,
	color.Black,
	color.RGBA{R: 255, G: 0, B: 0, A: 255},   // 1 Merah
	color.RGBA{R: 0, G: 255, B: 0, A: 255},   // 2 Hijau
	color.RGBA{R: 0, G: 0, B: 255, A: 255},   // 3 Biru
	color.RGBA{R: 255, G: 255, B: 0, A: 255}, // 4 Kuning
	color.RGBA{R: 255, G: 0, B: 255, A: 255}, // 5 Ungu
	color.RGBA{R: 0, G: 255, B: 255, A: 255}, // 6 Cyan
}

const (
	whiteIndex  = 0 // first color in palette
	blackIndex  = 1
	redIndex    = 2
	greenIndex  = 3
	blueIndex   = 4
	yellowIndex = 5
	purpleIndex = 6
	cyanIndex   = 7
)

func Lissajous(out io.Writer, c *string) error {
	const (
		res     = 0.0001
		size    = 100
		nframes = 64
		delay   = 8
	)

	var cycles float64

	number, err := strconv.ParseFloat(*c, 64)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	cycles = number

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	index_pallete := 1
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		t := 0.0
		index_pallete++
		for t < cycles*2*math.Pi {
			x := int(math.Sin(t)*float64(size) + 0.5)
			y := int(math.Sin(t*freq+phase)*float64(size) + 0.5)

			pixelX := size + x
			pixelY := size + y

			if pixelX >= 0 && pixelY >= 0 && pixelX < 2*size+1 && pixelY < 2*size+1 {
				img.SetColorIndex(pixelX, pixelY, uint8(index_pallete))
			}

			t += res
		}

		if index_pallete >= 7 {
			index_pallete = 0
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim)
	return nil
}
