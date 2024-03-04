package piscine

func Sqrt2(nb int) int {
	for i := 0; i*i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}
