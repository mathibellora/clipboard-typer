package main

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
)

func makeIcon() []byte {
	const size = 32
	img := image.NewRGBA(image.Rect(0, 0, size, size))

	bg := color.RGBA{30, 30, 50, 255}
	fg := color.RGBA{200, 150, 80, 255}

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			img.Set(x, y, bg)
		}
	}

	// keyboard outline
	for y := 8; y <= 24; y++ {
		for x := 3; x <= 28; x++ {
			if y == 8 || y == 24 || x == 3 || x == 28 {
				img.Set(x, y, fg)
			}
		}
	}

	// top row keys
	for _, kx := range []int{6, 11, 16, 21} {
		for y := 11; y <= 14; y++ {
			for x := kx; x <= kx+2; x++ {
				img.Set(x, y, fg)
			}
		}
	}

	// bottom row keys
	for _, kx := range []int{8, 14, 19} {
		for y := 17; y <= 20; y++ {
			for x := kx; x <= kx+2; x++ {
				img.Set(x, y, fg)
			}
		}
	}

	var buf bytes.Buffer
	png.Encode(&buf, img)
	return buf.Bytes()
}
