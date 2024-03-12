package piscine

func ForEach(f func(int), a []int) {
	for _, val := range a {
		f(val)
	}
}
