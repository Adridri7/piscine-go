package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for idx, arg := range args {
		for _, char := range arg {
			if char != '/' && char != '.' {
				z01.PrintRune(char)
			}
		}
		if idx != len(args)-1 {
			z01.PrintRune('\n')
		}
	}
}
