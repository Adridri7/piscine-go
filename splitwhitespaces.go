package piscine

func SplitWhiteSpaces(s string) (tab []string) {
	tmp := ""

	for _, char := range s {
		if char == ' ' {
			tab = append(tab, tmp)
			tmp = ""

		} else {
			tmp += string(char)
		}
	}
	if tmp != "" {
		tab = append(tab, tmp)
	}
	return tab
}
