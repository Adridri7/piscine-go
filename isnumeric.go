package piscine

func IsNumeric(s string) bool {
	for _, val := range s {
		if val < '0' || val > '9' {
			return false
		}
	}
	return true
}
