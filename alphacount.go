package piscine

func AlphaCount(s string) int {
	cpt := 0
	for _, val := range s {
		if val >= 'a' && val <= 'z' || val >= 'A' && val <= 'Z' {
			cpt++
		}
	}
	return cpt
}
