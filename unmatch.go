package piscine

func Unmatch(arr []int) int {
	var count int
	for _, el := range arr {
		count = 0
		for _, v := range arr {
			if v == el {
				count++
			}
		}
		if count%2 != 0 {
			return el
		}
	}
	return -1
}
