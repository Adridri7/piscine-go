package piscine

func TrimAtoi(s string) int {
	tmp := ""
	cpt := 0
	for _, val := range s {
		if val >= '0' && val <= '9' {
			tmp += string(val)
			cpt++
		} else if val == '-' {
			if cpt != 1 {
				tmp = "-" + tmp
			}
		}
	}

	return Atoi(tmp)
}
