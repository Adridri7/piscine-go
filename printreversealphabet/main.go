package main

import (
	"github.com/01-edu/z01"
)

func main() {
	var i rune
	for i = 122; i >= 97; i-- {
		z01.PrintRune(i)
	}

	z01.PrintRune('\n')
}
