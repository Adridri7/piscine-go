package piscine

func Split(s, sep string) (tab []string) {
	tmp_str := ""
	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			tab = append(tab, tmp_str)
			tmp_str = ""
			i += len(sep) - 1
		} else {
			tmp_str += string(s[i])
		}
	}
	tmp_str += string(s[len(s)-len(sep)+1:len(s)])
	tab = append(tab, tmp_str)
	return tab
}