package piscine

func Max(a []int) int {
	max := a[0]
	for _, val := range a {
		if val > max {
			max = val
		}
	}
	return max
}
