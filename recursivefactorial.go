package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb >= 0 {
		return 1
	} else if nb > 65 {
		return 0
	}

	return nb * RecursiveFactorial(nb-1)

}
