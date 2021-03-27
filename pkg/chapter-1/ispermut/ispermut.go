package ispermut

import ()

func countRunes(s string) map[rune]int {
	counts := map[rune]int{}
	for _, r := range s {
		counts[r]++
	}
	return counts
}

func IsPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1Counts := countRunes(s1)
	s2Counts := countRunes(s2)
	for r1, c1 := range s1Counts {
		c2 := s2Counts[r1]
		if c1 != c2 {
			return false
		}
	}
	return true
}

func IsPermutationNoAlloc(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for _, r1 := range s1 {
		r1CountInS2 := 0
		for _, r2 := range s2 {
			if r2 == r1 {
				r1CountInS2++
			}
		}
		if r1CountInS2 == 0 {
			return false
		}
		r1CountInS1 := 0
		for _, r1Inner := range s1 {
			if r1Inner == r1 {
				r1CountInS1++
			}
		}
		if r1CountInS1 != r1CountInS2 {
			return false
		}
	}

	return true
}
