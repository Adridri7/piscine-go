package piscine

func FirstRune(s string) rune {
	if len(s) >= 1 {
		return rune(s[0])
	}
	return '\n'
}
