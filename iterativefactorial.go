package piscine

func IterativeFactorial(nb int) int {
	if nb < 1 {
		return 0
	} else if nb > 65 {
		return 0
	}

	res := 1
	for i := 1; i <= nb; i++ {
		res *= i
	}
	return res
}
