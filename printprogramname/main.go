package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for _, char := range args[0] {
		if char != '/' && char != '.' {
			z01.PrintRune(char)
		}
	}
}
