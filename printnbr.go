package piscine

import (
	"github.com/01-edu/z01"
)

var r []rune

func PrintNbr(n int) {
	var n1 uint16
	if n < 0 {
		z01.PrintRune('-')
		n *= -1
	}

	if n == 0 {
		z01.PrintRune('0')
	}

	n1 = uint16(n)

	for n1 != 0 {
		r = append(r, rune(n1%10+'0'))
		n1 /= 10
	}

	for i := len(r) - 1; 0 <= i; i-- {
		z01.PrintRune(r[i])
	}
}
