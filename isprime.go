package piscine

func IsPrime(nb int) bool {
	tab := []int{}
	for i := 1; i <= nb; i++ {
		if nb%i == 0 {
			tab = append(tab, i)
		}
	}

	if len(tab) == 2 {
		cpt := 0
		for _, val := range tab {
			if val == 1 || val == nb {
				cpt++
			}
		}
		if cpt == 2 {
			return true
		} else {
			return false
		}
	}
	return false
}
