package piscine

func Capitalize(s string) string {
	tmp := ""
	res := ""
	for _, val := range s {
		if !(val >= 'A' && val <= 'Z') && !(val >= 'a' && val <= 'z') && !(val >= '0' && val <= '9') {

			res += maj(tmp)
			res += string(val)
			tmp = ""
		} else {
			tmp += string(val)
		}

	}
    res += maj(tmp)
	return res
}

func maj(s string) string {
	res := ""
	tab := []string{}
	for idx, char := range s {
		if idx == 0 && char >= 'a' && char <= 'z' {
			tab = append(tab, string(char-32))
		} else if idx == 0 && char >= 'A' && char <= 'Z' {
			tab = append(tab, string(char))
		} else {
			if char >= 'A' && char <= 'Z' {
				tab = append(tab, string(char+32))
			} else {
				tab = append(tab, string(char))
			}
		}
	}

	for _, val := range tab {
		res += string(val)
	}
	return res
}
