package piscine

func Sqrt(nb int) int {
	if nb == 0 || nb < 0 {
		return 0
	} else if nb == 1 {
		return 1
	}

	z := nb / 2
	for z*z-nb > 0 {
		z = (z + nb/z) / 2
	}

	if z*z != nb {
		return 0
	}

	return z
}
