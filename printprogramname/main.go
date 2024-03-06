package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	args := os.Args
	print(args[0])

}

func print(s string) {
	for _, char := range s {
		z01.PrintRune(char)
	}
}
