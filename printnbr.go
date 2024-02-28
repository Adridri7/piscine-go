package piscine

import (
	"github.com/01-edu/z01"
)

func MyItos(n int) string {
	var numberToString string
	var isNegative bool

	if n == 0 {
		return "0"
	}
	if n < 0 {
		isNegative = true
		n = -n
	}
	for n > 0 {
		var tempN int = n % 10
		numberToString = string(rune(tempN+'0')) + numberToString
		n /= 10
	}
	if isNegative {
		numberToString = "-" + numberToString
	}

	return numberToString
}

func PrintNbr(n int) {
	for _, val := range MyItos(n) {
		z01.PrintRune(val)
	}
}
