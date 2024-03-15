package piscine

func Unmatch(a []int) int {
	numCount := make(map[int]int)

	for _, num := range a {
		numCount[num]++
	}

	for num, count := range numCount {
		if count%2 != 0 {
			return num
		}
	}
	return -1
}
