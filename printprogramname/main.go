package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	printstrRune(args[0])
}

func printstrRune(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}
