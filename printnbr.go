package piscine

import (
	"github.com/01-edu/z01"
)

var r []rune

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
		var n1 uint
		if n < 0 {
			z01.PrintRune('-')
			n *= -1
		}

		if n == 0 {
			z01.PrintRune('0')
		}

		n1 = uint(n)

		for {
			if n1 == 0 {
				break
			}
			r = append(r, rune(n1%10+'0'))
			n1 /= 10
		}

		for i := len(r) - 1; 0 <= i; i-- {
			z01.PrintRune(r[i])
		}
	}
}
