package piscine

func NRune(s string, n int) rune {
	r := []rune(s)

	if n > len(r)-1 {
		return 0
	} else if n < 0 {
		if len(r)-1+n < 0 {
			return 0
		}
		n = len(r) + n + 1
	}

	return r[n-1]
}
