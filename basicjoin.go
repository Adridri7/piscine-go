package piscine

func BasicJoin(elems []string) string {
	res := ""
	for _, val := range elems {
		res += string(val)
	}
	return res
}
