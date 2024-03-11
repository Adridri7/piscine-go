package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printStr("File name missing")
		return
	}
	for _, arg := range args {
		file, err := os.ReadFile(arg)
		if err != nil {
			printStr(err.Error())
			return
		}
		printStr(string(file))
	}
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
