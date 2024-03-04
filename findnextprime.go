package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) {
		return nb
	} else {
		for val := nb; ; val++ {
			if IsPrime(val) {
				return val
			}
		}
	}
}
