package runes

func InWord(word string, r rune) bool {
	for _, wr := range word {
		if wr == r {
			return true
		}
	}
	return false
}

func Match(r1, r2 byte) bool {
	if r2 == '?' {
		return true
	}
	return r1 == r2
}

func DoesntMatch(r1, r2 byte) bool {
	if r2 == '?' {
		return false
	}
	return r1 == r2
}
