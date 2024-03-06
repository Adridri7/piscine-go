package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for _, arg := range args {
		for _, char := range arg {
			if char != '/' && char != '.' {
				z01.PrintRune(char)
			}
		}
		z01.PrintRune('\n')
	}
}
