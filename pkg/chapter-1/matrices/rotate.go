package matrices

import (
	"image"
	"image/color"
)

func matrixToImage(yxM [][]uint32) *image.CMYK {
	x0, y0 := 0, 0
	x1, y1 := len(yxM[0]), len(yxM)
	img := image.NewCMYK(image.Rect(x0, y0, x1, y1))
	for row := 0; row < len(yxM); row++ {
		for col := 0; col < len(yxM[0]); col++ {
			colorEncoded := yxM[row][col]
			colorDecoded := color.CMYK{
				C: uint8(colorEncoded),
				M: uint8(colorEncoded >> 8),
				Y: uint8(colorEncoded >> 16),
				K: uint8(colorEncoded >> 24),
			}
			img.SetCMYK(col, row, colorDecoded)
		}
	}
	return img
}

// TODO: move channel out of the API into a global var
func Rotate(yxM [][]uint32, ch chan image.Image) {
	if len(yxM) == 0 {
		panic("an empty matrix passed to Rotate")
	}
	if len(yxM[0]) != len(yxM) {
		panic("a non-square matrix passed to Rotate")
	}

	if ch != nil {
		ch <- matrixToImage(yxM)
	}

	rowLen := len(yxM)
	for y0 := 0; y0 < (rowLen / 2); y0++ {
		for x0 := y0; x0 < (rowLen - y0 - 1); x0++ {
			tmp := yxM[y0][x0]

			y3 := rowLen - 1 - x0
			x3 := y0
			yxM[y0][x0] = yxM[y3][x3]

			y2 := rowLen - 1 - y0
			x2 := rowLen - 1 - x0
			yxM[y3][x3] = yxM[y2][x2]

			y1 := x0
			x1 := rowLen - 1 - y0
			yxM[y2][x2] = yxM[y1][x1]

			yxM[y1][x1] = tmp

			if ch != nil {
				img := matrixToImage(yxM)
				img.SetCMYK(x0, y0, color.CMYK{C: 0xFF})
				img.SetCMYK(x1, y1, color.CMYK{M: 0xFF})
				img.SetCMYK(x2, y2, color.CMYK{Y: 0xFF})
				img.SetCMYK(x3, y3, color.CMYK{K: 0xFF})
				ch <- img
			}
		}
	}

	if ch != nil {
		ch <- matrixToImage(yxM)
		close(ch)
	}
}
