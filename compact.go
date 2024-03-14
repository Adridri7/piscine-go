package piscine

func Compact(ptr *[]string) int {
	cpt := 0
	for _, val := range *ptr {
		if val != "" {
			cpt++
		}
	}
	return cpt
}
