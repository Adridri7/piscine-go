package piscine

func IsAlpha(s string) bool {
	for _, val := range s {
		if !(val >= 'A' && val <= 'z') && !(val >= '0' && val <= '9') {
			return false
		}
	}
	return true
}
