package matrices

func ZeroCrossForZeroElementAlloc(yxM [][]uint32) {
	if len(yxM) == 0 {
		panic("an empty matrix passed to Zero")
	}

	rowsToZero := map[int]bool{}
	colsToZero := map[int]bool{}
	for y := 0; y < len(yxM); y++ {
		for x := 0; x < len(yxM[0]); x++ {
			if yxM[y][x] == 0 {
				rowsToZero[y] = true
				colsToZero[x] = true
			}
		}
	}

	for y := 0; y < len(yxM); y++ {
		for x := 0; x < len(yxM[0]); x++ {
			if rowsToZero[y] || colsToZero[x] {
				yxM[y][x] = 0
			}
		}
	}
}

func ZeroCrossForZeroElement(yxM [][]uint32) {
	if len(yxM) == 0 {
		panic("an empty matrix passed to Zero")
	}

	rowsWithSpecialData := uint32(0)
	colsWithSpecialData := uint32(0)
	for y := 0; y < len(yxM); y++ {
		for x := 0; x < len(yxM[0]); x++ {
			if yxM[y][x] == 0 {
				rowsWithSpecialData = rowsWithSpecialData | (1 << y)
				colsWithSpecialData = colsWithSpecialData | (1 << x)
				yxM[y][0] = yxM[y][0] | (1 << x)
				yxM[0][x] = yxM[0][x] | (1 << y)
			}
		}
	}

	for y := 0; y < len(yxM); y++ {
		for x := 0; x < len(yxM[0]); x++ {
			if (rowsWithSpecialData&(1<<y)) > 0 && (yxM[y][0]&(1<<x)) > 0 {
				yxM[y][x] = 0
			}
			if (colsWithSpecialData&(1<<x)) > 0 && (yxM[0][x]&(1<<y)) > 0 {
				yxM[y][x] = 0
			}
		}
	}

	for y := 0; y < len(yxM); y++ {
		for x := 0; x < len(yxM[0]); x++ {
			if (rowsWithSpecialData&(1<<y)) > 0 || (colsWithSpecialData&(1<<x)) > 0 {
				yxM[y][x] = 0
			}
		}
	}
}
