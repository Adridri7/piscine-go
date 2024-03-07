package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, str := range a {
		printstr(str)
		z01.PrintRune('\n')
	}
}

func printstr(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}
