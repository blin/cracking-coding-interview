package compress

import (
	"strconv"
	"strings"
)

func Compress(s string) string {
	var compressed strings.Builder

	var currentRune rune
	var count int
	for i, r := range s {
		if i == 0 {
			currentRune = r
			count = 1
			continue
		}

		if r == currentRune {
			count++
			continue
		}
		compressed.WriteRune(currentRune)
		compressed.WriteString(strconv.Itoa(count))
		currentRune = r
		count = 1
	}
	compressed.WriteRune(currentRune)
	compressed.WriteString(strconv.Itoa(count))

	if compressed.Len() > len(s) {
		return s
	}
	return compressed.String()
}
