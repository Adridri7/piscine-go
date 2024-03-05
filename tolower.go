package piscine

func ToLower(s string) string {
	tmp := ""
	for _, val := range s {
		if val >= 'A' && val <= 'Z' {
			tmp += string(val + 32)
		} else {
			tmp += string(val)
		}
	}
	return tmp
}
