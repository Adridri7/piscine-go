package piscine

func Atoi(str string) int {
	result := 0
	sign := 1

    if len(str) == 0 {
		return 0 // Retourner 0 si la chaîne est vide
	}

    if len(str) == 1 {
        digit := int(str[0] - '0')
		result = result*10 + digit
		return 0 // Retourner 0 si la chaîne est vide
	}

  

	// Gérer le signe
	if str[0] == '-' {
		sign = -1
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}

	for _, char := range str {
		if char < '0' || char > '9' {
			return 0
		}
		digit := int(char - '0')
		result = result*10 + digit
	}

	return result * sign
}
