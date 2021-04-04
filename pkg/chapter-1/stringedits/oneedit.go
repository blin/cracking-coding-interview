package stringedits

func isOneRuneChangeAway(s1, s2 string) bool {
	if s1 == s2 {
		return false
	}
	edits := 0
	for i := 0; i < len(s1); i++ {
		r1 := s1[i]
		r2 := s2[i]
		if r1 != r2 {
			edits++
		}
	}
	return edits == 1
}

func isOneAdditionAway(s1, s2 string) bool {
	for i := 0; i < len(s1); i++ {
		r1 := s1[i]
		r2 := s2[i]
		if r1 == r2 {
			continue
		}
		return s1[i:] == s2[i+1:]
	}
	return true
}

func IsOneEditAway(s1, s2 string) bool {
	lenDiff := len(s1) - len(s2)
	if lenDiff > 1 || lenDiff < -1 {
		return false
	}

	if len(s1) == len(s2) {
		return isOneRuneChangeAway(s1, s2)
	}

	if len(s2) > len(s1) {
		return isOneAdditionAway(s1, s2)
	}

	return isOneAdditionAway(s2, s1)

}
