package piscine

func MakeRange(min, max int) []int {
	tab := make([]int, max)
	for i := min; i <= max; i++ {
		tab[i] = i + 1
	}
	return tab
}
