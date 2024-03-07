package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for idx, str := range a {
		printstr(str)
		if idx != len(a)-1 {
			z01.PrintRune('\n')
		}

	}
}

func printstr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}
