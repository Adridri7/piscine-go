package piscine

func TrimAtoi(s string) int {
	tmp := ""
	i := -1
	for _, val := range s {
		if val >= '0' && val <= '9' {
			i++
			tmp += string(val)

		} else if val == '-' {
			if i == -1 {
				tmp = "-" + tmp
			}
		}
	}
	return Atoi(tmp)
}
