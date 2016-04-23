package main

import (
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"
)

func main() {
	height := 480
	width := 640
	centerX := width / 2
	centerY := height / 2
	radius := width / 4
	red := color.RGBA{255, 0, 0, 255}
	white := color.RGBA{255, 255, 255, 255}
	image := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			distanceFromCenter := int(math.Sqrt(
				math.Pow(float64(y-centerY), 2) + math.Pow(float64(x-centerX), 2)))

			if distanceFromCenter < radius {
				image.Set(x, y, red)
			} else {
				image.Set(x, y, white)
			}
		}
	}

	file, _ := os.Create("japan.jpeg")
	jpeg.Encode(file, image, nil)
}
