package piscine

func FirstRune(s string) rune {
	if len(s) >= 1 {
		for _, char := range s {
			return char
		}
	}
	return rune('\n')
}
