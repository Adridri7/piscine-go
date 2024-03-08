package piscine

func AtoiBase(s string, base string) int {
	res := 0
	if !isValidBase(base) {
		return res
	}
	cpt := len(s) - 1
	idx := 0
	for _, char := range s {
		for ind, c := range base {
			if char == c {
				idx = ind
			}
		}
		res += idx * RecursivePower(len(base), cpt)
		cpt--
	}
	return res
}
