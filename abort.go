package piscine

func Abort(a, b, c, d, e int) int {
	tab := []int{a, b, c, d, e}
	sortedTab := sortIntegerTable(tab)

	return sortedTab[2]
}
