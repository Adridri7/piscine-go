package piscine

func fonction(s string, sep string) []string {
	//création d'un tableau
	tab := []string{}
	startindex := 0
	Lsep := len(sep)

	//boucle jusqu'a len - len du sep pour eviter runtime
	for i := 0; i < len(s)-Lsep+1; i++ {

		//BANGER
		if s[i:i+Lsep] == sep {
			//CLASSICO
			tab = append(tab, s[startindex:i])
			startindex = i + Lsep
			i = i + Lsep - 1
		}
	}
	tab = append(tab, s[startindex:])
	return tab
}
