package urlify

func URLify(s []rune) []rune {
	spaces := 0
	for _, r := range s {
		if r == ' ' {
			spaces++
		}
	}

	urlified := make([]rune, len(s)+(spaces*2))

	j := cap(urlified) - 1
	for i := len(s) - 1; i >= 0; i-- {
		r := s[i]

		if r != ' ' {
			urlified[j] = r
			j--
			continue
		}
		urlified[j-2] = '%'
		urlified[j-1] = '2'
		urlified[j] = '0'
		j = j - 3
	}

	return urlified
}
