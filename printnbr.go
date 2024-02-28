package piscine

import (
	"github.com/01-edu/z01"
)

func MyItos(n int) string {
	var numberToString string
	if n == 0 {
		return "0"
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	for n != 0 {
		numberToString = string(rune(n%10+'0')) + numberToString
		n /= 10
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
