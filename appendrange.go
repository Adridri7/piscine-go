package piscine

func AppendRange(min, max int) (tab []int) {
	for i := min; i < max; i++ {
		tab = append(tab, i)
	}
	return tab
}
