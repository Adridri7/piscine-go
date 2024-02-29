package piscine


func StrLen(s string) int {
	count := 0
	for idx := range s {
		count = idx
	}

	return count
}