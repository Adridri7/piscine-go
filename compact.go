package piscine

func Compact(ptr *[]string) int {
	tab := []string{}
	cpt := 0
	for _, val := range *ptr {
		if val != "" {
			cpt++
			tab = append(tab, val)
		}
	}
	*ptr = tab
	return cpt
}
