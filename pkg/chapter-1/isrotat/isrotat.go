package isrotat

import "strings"

// Had no clue how to use Contains until third hint :(
func IsRotation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	return strings.Contains(s2+s2, s1)
}

func IsRotationAlternativeImplementation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i2 := 0; i2 < len(s2); i2++ {
		if s2[i2] != s1[0] {
			continue
		}
		s2Rotated := s2[i2:] + s2[:i2]
		if s1 == s2Rotated {
			return true
		}
	}
	return false
}
