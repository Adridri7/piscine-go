package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i := 0; i <= len(a)-1; i++ {
		for j := i + 1; j <= len(a)-1; j++ {
			if check(a[i], a[j]) == -1 {
				return false
			}
		}
	}
	return true
}

func check(a, b int) int {
	if a > b {
		return 1
	} else if a == b {
		return 0
	} else {
		return -1
	}
}
