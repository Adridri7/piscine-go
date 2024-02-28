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
	n2 := int64(n)
	println(n)
	for n2 < 0 {
		var tempN int64 = n2 % 10
		numberToString = string(rune(tempN+'0')) + numberToString
		n2 /= 10
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
