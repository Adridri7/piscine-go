package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	t := sortIntegerTable(a)
	for i := 0; i <= len(a)-1; i++ {
		if t[i] != a[i] {
			return false
		}
	}
	return true
}

func sortIntegerTable(table []int) []int {
	tab := table
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab); j++ {
			if tab[j] > table[i] {
				Swap(&tab[i], &tab[j])
			}
		}
	}
	return tab
}
