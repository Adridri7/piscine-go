package piscine

func BasicAtoi(s string) int {
	result := 0

	for _, char := range s {
		digit := int(char - '0')
		result = result*10 + digit
	}

	return result
}
