package piscine

import (
	"strconv"

	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	for _, char := range Int_to_str(n) {
		z01.PrintRune(char)
	}
}

func Int_to_str(n int) string {
	return strconv.Itoa(n)
}
