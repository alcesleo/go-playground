package main

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"math"
	"os"
)

var height = 100
var width = 100
var centerX = width / 2
var centerY = height / 2
var color1 = color.RGBA{0, 255, 0, 255}
var color2 = color.RGBA{255, 0, 130, 255}

func main() {
	var images []*image.Paletted
	var delays []int
	steps := height / 2

	for step := 0; step < steps; step++ {
		img := drawCircle(step*2, color1, color2)
		images = append(images, img)
		delays = append(delays, 0)
	}

	for step := 0; step < steps; step++ {
		img := drawCircle(step*2, color2, color1)
		images = append(images, img)
		delays = append(delays, 0)
	}

	file, _ := os.Create("gof.gif")
	defer file.Close()
	gif.EncodeAll(file, &gif.GIF{
		Image: images,
		Delay: delays,
	})
}

func drawCircle(radius int, fg, bg color.RGBA) *image.Paletted {
	image := image.NewPaletted(image.Rect(0, 0, width, height), palette.WebSafe)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			distanceFromCenter := int(math.Sqrt(
				math.Pow(float64(y-centerY), 2) + math.Pow(float64(x-centerX), 2)))

			if distanceFromCenter < radius {
				image.Set(x, y, fg)
			} else {
				image.Set(x, y, bg)
			}
		}
	}

	return image
}
