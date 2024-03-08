package piscine

func Split(s string, sep string) []string {
	tab := []string{}
	startindex := 0
	Lsep := len(sep)
	for i := 0; i < len(s)-Lsep+1; i++ {
		if s[i:i+Lsep] == sep {
			tab = append(tab, s[startindex:i])
			startindex = i + Lsep
			i = i + Lsep - 1
		}
	}
	tab = append(tab, s[startindex:])
	return tab
}
