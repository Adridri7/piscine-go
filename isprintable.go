package piscine

func IsPrintable(s string) bool {
	for _, val := range s {
		if val < 22 || val > 127 {
			return false
		}
	}
	return true
}
