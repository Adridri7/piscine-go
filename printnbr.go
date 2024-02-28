package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	for _, val := range intToString(n) {
		z01.PrintRune(val)
	}
}

func intToString(num int) string {
	if num == 0 {
		return "0"
	}

	var result []rune

	isNegative := false
	if num < 0 {
		isNegative = true
		num = -num
	}

	for num != 0 {
		digit := num % 10
		result = append(result, rune(digit+'0'))
		num /= 10
	}

	if isNegative {
		result = append(result, '-')
	}

	// Inverse la chaîne résultante
	length := len(result)
	for i := 0; i < length/2; i++ {
		result[i], result[length-i-1] = result[length-i-1], result[i]
	}

	return string(result)
}
