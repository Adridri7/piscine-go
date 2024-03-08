package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) string {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
	} else {
		str := ""
		var n uint
		if nbr == 0 {
			i := nbr % len(base)
			str = string(base[i]) + str
		}
		if nbr < 0 {
			n = uint(-nbr)
			for n != 0 {
				i := n % uint(len(base))
				str = string(base[i]) + str
				n /= uint(len(base))
			}
			str = "-" + str
		} else {
			for nbr != 0 {
				i := nbr % len(base)
				str = string(base[i]) + str
				nbr /= len(base)
			}
		}
		return str
		//for _, char := range str {
		//	z01.PrintRune(char)
		//}
	}
	return ""

}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for _, val := range base {
		if val == '-' || val == '+' {
			return false
		} else if !contain1(base, val) {
			return false
		}
	}
	return true
}

func contain1(base string, r rune) bool {
	cpt := 0
	for _, val := range base {
		if val == r {
			cpt++
		}
	}
	return cpt == 1
}
