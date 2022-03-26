package pixel

import (
	"testing"
)

func TestPixel_Is(t *testing.T) {
	pixel := Pixel{R: 128, G: 256, B: 512, A: 64}

	tests := []struct {
		name       string
		r, g, b, a uint32
		expected   bool
	}{
		{"all values match", 128, 256, 512, 64, true},
		{"one value differs", 128, 128, 512, 64, false},
		{"all values differ", 1, 2, 3, 4, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := pixel.Is(tt.r, tt.g, tt.b, tt.a)
			if actual != tt.expected {
				t.Fatalf("expected %t, got %t", tt.expected, actual)
			}
		})
	}
}
