package piscine

func ActiveBits(n int) int {
	cpt := 0
	for n != 0 {
		if n%2 == 1 {
			cpt++
		}
		n /= 2
	}
	return cpt
}
