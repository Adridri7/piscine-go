package piscine

func Atoi(s string) int {
	result := 0
	sign := 1

	if s[0] == '-' {
		sign = -1
	}

	for _, char := range s {
		if char < '0' || char > '9' {
			return 0
		}
		digit := int(char - '0')
		result = result*10 + digit
	}

	return result * sign
}
