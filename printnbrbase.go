package piscine

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	if !isValidBase(base) {
		z01.PrintRune('N')
		z01.PrintRune('V')
	} else {
		str := ""
		is_negative := false
		if nbr < 0 {
			is_negative = true
			nbr = -nbr
		}
		if nbr == 0 {
			i := nbr % len(base)
			str = string(base[i]) + str
		} else {
			for nbr != 0 {
				i := nbr % len(base)
				if i >= 0 {
					str = string(base[i]) + str
				}
				nbr /= len(base)
			}
		}
		if is_negative {
			str = "-" + str
		}
		for _, char := range str {
			z01.PrintRune(char)
		}
	}
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
