package piscine

func NRune(s string, n int) rune {
	if n < 0 {
		return 0
	}
	for idx, char := range s {
		if idx == n-1 {
			return rune(char)
		}
	}
	return 0
}
