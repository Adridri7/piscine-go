package main

import (
	"github.com/01-edu/z01"
)

func main() {
	var i rune
	for i = 48; i <= 57; i++ {
		z01.PrintRune(i)
	}

	z01.PrintRune('\n')
}
