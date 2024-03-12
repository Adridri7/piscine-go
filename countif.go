package piscine

func CountIf(f func(string) bool, tab []string) int {
	cpt := 0
	for _, str := range tab {
		if f(str) {
			cpt++
		}
	}
	return cpt
}
