package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) >= 1 {
		if args[0] == "--upper" {
			for i := 1; i <= len(args)-1; i++ {
				idx := atoi(args[i])
				if idx < 0 || idx > 26 {
					z01.PrintRune(' ')
				} else {
					z01.PrintRune(rune('A' - 1 + idx))
				}
			}
		} else {
			for i := 0; i <= len(args)-1; i++ {
				idx := atoi(args[i])
				if idx < 0 || idx > 26 {
					z01.PrintRune(' ')
				} else {
					z01.PrintRune(rune('a' - 1 + idx))
				}
			}
		}
	}
}

func atoi(s string) int {
	result := 0

	for _, char := range s {
		digit := int(char - '0')
		result = result*10 + digit
	}

	return result
}
