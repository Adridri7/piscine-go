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

	if n == -9223372036854775808 {
		z01.PrintRune('-')
		z01.PrintRune('9')
		z01.PrintRune('2')
		z01.PrintRune('2')
		z01.PrintRune('3')
		z01.PrintRune('3')
		z01.PrintRune('7')
		z01.PrintRune('2')
		z01.PrintRune('0')
		z01.PrintRune('3')
		z01.PrintRune('6')
		z01.PrintRune('8')
		z01.PrintRune('5')
		z01.PrintRune('4')
		z01.PrintRune('7')
		z01.PrintRune('7')
		z01.PrintRune('5')
		z01.PrintRune('8')
		z01.PrintRune('0')
		z01.PrintRune('8')
	} else {
		for _, val := range MyItos(n) {
			z01.PrintRune(val)
		}
	}
}
