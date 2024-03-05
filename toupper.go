package piscine

func ToUpper(s string) string {
	tmp := ""
	for _, val := range s {
		if val >= 'a' && val <= 'z' {
			tmp += string(val - 32)
		} else {
			tmp += string(val)
		}
	}
	return tmp
}
