package piscine

func NRune(s string, n int) rune {
	if n > len(s)-1 || -n > len(s) {
		return 0
	}
	if n < 0 {
		return rune(s[len(s)+n])
	}
	return rune(s[n-1])
}
