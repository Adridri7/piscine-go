package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	args := os.Args
	sortAlpha(args)

	for i := len(args) - 1; i > 0; i-- {
		for _, char := range args[i] {
			if char != '/' && char != '.' {
				z01.PrintRune(char)
			}
		}
		z01.PrintRune('\n')
	}
}

func sortAlpha(tab []string) {
	for i := 1; i <= len(tab)-1; i++ {
		for j := i + 1; i <= len(tab)-1; i++ {
			if tab[i] < tab[j] {
				swapstr(&tab[i], &tab[j])
			}
		}

	}
}

func swapstr(a *string, b *string) {
	c := *a
	*a = *b
	*b = c
}
