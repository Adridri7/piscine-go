package piscine

func IsUpper(s string) bool {
	for _, val := range s {
		if val < 'A' || val > 'Z' {
			return false
		}
	}
	return true
}
