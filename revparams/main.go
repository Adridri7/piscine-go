package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args

	for i := len(args) - 1; i > 0; i-- {
		if i != 0 {
			for _, char := range args[i] {
				if char != '/' && char != '.' {
					z01.PrintRune(char)
				}
			}
			z01.PrintRune('\n')
		}
	}
}
