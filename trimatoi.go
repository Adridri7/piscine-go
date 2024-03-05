package piscine

func TrimAtoi(s string) int {
	tmp := ""
	var ind int = -1
	for idx, val := range s {
		if val >= '0' && val <= '9' {
			tmp += string(val)
			ind = idx
		} else if val == '-' {
			if ind == -1 {
				tmp = "-" + tmp
			}
		}
	}

	return Atoi(tmp)
}
