package piscine

func Join(strs []string, sep string) string {
	res := ""
	for idx, val := range strs {
		res += string(val)
		if idx != len(strs)-1 {
			res += sep
		}
	}
	return res
}
