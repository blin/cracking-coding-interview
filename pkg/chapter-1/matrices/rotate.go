package matrices

import (
	"fmt"
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

func computeSquarePosition(sideSize, depth, idx int) (y int, x int) {
	sideSizeAtDepth := sideSize - (depth * 2)

	maxTop := sideSizeAtDepth - 1
	maxRight := maxTop + sideSizeAtDepth - 1
	maxBottom := maxRight + sideSizeAtDepth - 1
	maxLeft := maxBottom + sideSizeAtDepth - 2

	maxIdx := maxLeft
	if idx < 0 || idx > maxLeft {
		panic(fmt.Errorf("got idx=%d , want idx within [0, %d]", idx, maxIdx))
	}

	farthestAtDepth := sideSize - depth - 1

	if idx >= 0 && idx <= maxTop {
		y = depth
		x = depth + idx
		return
	}

	if idx > maxTop && idx <= maxRight {
		y = depth + (idx - maxTop)
		x = farthestAtDepth
		return
	}

	if idx > maxRight && idx <= maxBottom {
		y = farthestAtDepth
		x = farthestAtDepth - (idx - maxRight)
		return
	}

	y = farthestAtDepth - (idx - maxBottom)
	x = depth
	return
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

	sideSize := len(yxM)
	for depth := 0; depth < (sideSize / 2); depth++ {
		sideSizeAtDepth := sideSize - (depth * 2)
		for i := 0; i < (sideSizeAtDepth - 1); i++ {
			y0, x0 := computeSquarePosition(sideSize, depth, i+((sideSizeAtDepth-1)*0))
			y1, x1 := computeSquarePosition(sideSize, depth, i+((sideSizeAtDepth-1)*1))
			y2, x2 := computeSquarePosition(sideSize, depth, i+((sideSizeAtDepth-1)*2))
			y3, x3 := computeSquarePosition(sideSize, depth, i+((sideSizeAtDepth-1)*3))

			tmp := yxM[y0][x0]
			yxM[y0][x0] = yxM[y3][x3]
			yxM[y3][x3] = yxM[y2][x2]
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
