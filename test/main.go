package main

import (
	"fmt"
	"piscine"

	"github.com/01-edu/z01"
)

func main() {
	//piscine.PrintNbrBase(125, "0123456789")
	//z01.PrintRune('\n')
	//piscine.PrintNbrBase(-125, "01")
	//z01.PrintRune('\n')
	//piscine.PrintNbrBase(125, "0123456789ABCDEF")
	//z01.PrintRune('\n')
	//piscine.PrintNbrBase(-125, "choumi")
	//z01.PrintRune('\n')
	//piscine.PrintNbrBase(-9223372036854775808, "0123456789")
	fmt.Println(piscine.BasicAtoi2("12345"))
	fmt.Println(piscine.BasicAtoi2("0000000012345"))
	fmt.Println(piscine.BasicAtoi2("012 345"))
	fmt.Println(piscine.BasicAtoi2("Hello World!"))
	z01.PrintRune('\n')
}
