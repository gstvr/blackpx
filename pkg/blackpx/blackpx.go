package blackpx

import (
	"image"

	"github.com/gstvr/blackpx/pkg/pixel"
)

func CountTrueBlack(img image.Image) (blackPixels int) {
	return Count(img, pixel.TrueBlack)
}

func Count(img image.Image, pixel pixel.Pixel) (pixels int) {
	bounds := img.Bounds()
	var r, g, b, a uint32

	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			r, g, b, a = img.At(x, y).RGBA()
			if pixel.Is(r, g, b, a) {
				pixels++
			}

		}
	}
	return pixels
}
