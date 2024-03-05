package piscine

func IsLower(s string) bool {
	for _, val := range s {
		if val < 'a' || val > 'z' {
			return false
		}
	}
	return true
}
