package piscine

func Join(strs []string, sep string) string {
	res := ""
	for _, val := range strs {
		res += string(val)
		res += sep
	}
	return res
}
