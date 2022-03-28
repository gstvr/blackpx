package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/gstvr/blackpx/pkg/blackpx"
)

func main() {
	if len(os.Args) != 2 {
		exit("no filename provided")
	}

	pixels, err := countTrueBlackPixels(os.Args[1])
	if err != nil {
		exit(err.Error())
	}

	fmt.Fprintln(os.Stdout, pixels)
}

func countTrueBlackPixels(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("could not read file: %v", err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		return 0, fmt.Errorf("could not decode image: %v", err)
	}

	return blackpx.CountTrueBlack(img), nil
}

func exit(message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(1)
}
