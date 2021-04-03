package palindromes

func IsPalindromePermutation(s string) bool {
	runes := map[rune]int{}
	for _, r := range s {
		runes[r]++
	}
	delete(runes, ' ')

	oddRunesCount := 0
	for _, c := range runes {
		if c%2 == 0 {
			continue
		}
		oddRunesCount++
	}

	return (oddRunesCount == 1 || oddRunesCount == 0)
}
