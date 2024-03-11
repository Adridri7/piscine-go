package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	printStr(points.x, points.y)
}

func printStr(x, y int) {
	s := "x = " + Itoa(x) + ", y = " + Itoa(y)
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func Itoa(n int) (res string) {
	if n == 0 {
		res = "0"
	}
	if n < 0 {
		val := uint(-n)
		for val != 0 {
			res = string(rune(val%10+'0')) + res
			val /= 10
		}
		res = "-" + res
	} else {
		for n != 0 {
			res = string(rune(n%10+'0')) + res
			n /= 10
		}
	}
	return res
}
