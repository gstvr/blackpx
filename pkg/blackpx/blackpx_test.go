package blackpx

import (
	"image"
	_ "image/png"
	"os"
	"testing"

	"github.com/gstvr/blackpx/pkg/pixel"
)

const examplesDir = "../../test/images/"

func TestCount(t *testing.T) {
	whitePixel := pixel.Pixel{R: pixel.MaxAlpha, G: pixel.MaxAlpha, B: pixel.MaxAlpha, A: pixel.MaxAlpha}
	yellowPixel := pixel.Pixel{R: pixel.MaxAlpha, G: (pixel.MaxAlpha * 198) / 255, B: (pixel.MaxAlpha * 87) / 255, A: pixel.MaxAlpha}

	tests := []struct {
		name, filename string
		pixel          pixel.Pixel
		expected       int
	}{
		{"1 black pixel", "1-black-px.png", pixel.TrueBlack, 1},
		{"all but one black pixels", "1-non-black-px.png", pixel.TrueBlack, 100*100 - 1},
		{"1 white pixel", "1-non-black-px.png", whitePixel, 1},
		{"all but 6 black pixels", "6-non-black-pixels.png", pixel.TrueBlack, 100*100 - 6},
		{"all black pixels", "all-black.png", pixel.TrueBlack, 100 * 100},
		{"no black pixels", "no-black-pixels.png", pixel.TrueBlack, 0},
		{"all yellow pixels", "no-black-pixels.png", yellowPixel, 100 * 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			img := readImage(t, tt.filename)
			actual := Count(img, tt.pixel)
			if tt.expected != actual {
				t.Fatalf("expected %d, got %d", tt.expected, actual)
			}
		})
	}
}

func TestCountTrueBlack(t *testing.T) {
	img := readImage(t, "1-non-black-px.png")

	expected := 100*100 - 1
	actual := CountTrueBlack(img)
	if expected != actual {
		t.Fatalf("expected %d, got %d", expected, actual)
	}
}

func readImage(t testing.TB, filename string) image.Image {
	t.Helper()

	file, err := os.Open(examplesDir + filename)
	if err != nil {
		t.Fatalf("could not open example file [%s]: %v", filename, err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		t.Fatalf("could not decode image: %v", err)
	}

	return img
}
