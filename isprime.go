package piscine

func IsPrime(nb int) bool {
	tab := []int{}
	if nb == 2 { // Seul parmis les nombres pairs
		return true
	}
	if nb%2 == 0 {
		return false
	}
	for i := 3; i <= nb; i += 2 { // réduction par 2 du nombre d'itérations
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
