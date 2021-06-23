package radixsort

import (
	"math"
)

func findMaxWidth(xs []int) int {
	maxWidth := 1
	for _, x := range xs {
		w := int(math.Ceil(math.Log10(float64(x))))
		if w > maxWidth {
			maxWidth = w
		}
	}
	return maxWidth
}

func getNthDigit(x, n int) int {
	return int(math.Mod((float64(x) / math.Pow10(n-1)), 10))
}

func Sort(xs []int) {
	var buckets [10][]int
	maxWidth := findMaxWidth(xs)
	for i := 1; i <= maxWidth; i++ {
		for _, x := range xs {
			d := getNthDigit(x, i)
			buckets[d] = append(buckets[d], x)
		}

		j := 0
		for d := 0; d < 10; d++ {
			for _, x := range buckets[d] {
				xs[j] = x
				j++
			}
			buckets[d] = nil
		}
	}
}
