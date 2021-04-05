package main

import (
	"fmt"
	"image"
	"image/png"
	"math/rand"
	"os"
	"path"

	"github.com/blin/cracking-coding-interview/pkg/chapter-1/matrices"
)

func squareMatrix(size int) [][]uint32 {
	yxM := make([][]uint32, size)
	for row := 0; row < len(yxM); row++ {
		yxM[row] = make([]uint32, size)
		color := rand.Uint32()
		for col := 0; col < len(yxM[row]); col++ {
			yxM[row][col] = color
		}
	}

	return yxM
}

func main() {
	imagesDir := "/tmp/images/"
	if _, err := os.Stat(imagesDir); os.IsNotExist(err) {
		os.Mkdir(imagesDir, 0755)
	} else if err != nil {
		fmt.Printf("failed to stat %s : %v\n", imagesDir, err)
		os.Exit(2)
	}

	yxM := squareMatrix(50)
	imageCh := make(chan image.Image)

	go matrices.Rotate(yxM, imageCh)

	i := 0
	var pngEnc png.Encoder
	for img := range imageCh {
		fn := path.Join(imagesDir, fmt.Sprintf("%05d.png", i))
		f, err := os.Create(fn)
		if err != nil {
			fmt.Printf("failed to create file %s: %v\n", fn, err)
			os.Exit(2)
		}

		err = pngEnc.Encode(f, img)
		if err != nil {
			fmt.Printf("failed to encode an image: %v\n", err)
			os.Exit(2)

		}

		err = f.Close()
		if err != nil {
			fmt.Printf("failed to close file %s: %v\n", fn, err)
			os.Exit(2)

		}

		i++
	}
}
