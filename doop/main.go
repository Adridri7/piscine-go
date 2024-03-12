package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		return
	}

	a, _ := strconv.Atoi(args[0])
	b, _ := strconv.Atoi(args[2])

	calc(a, b, args[1])
}

func calc(a, b int, operator string) {
	if a+b > 9223372036854775807 || a+b < -9223372036854775808 || a-b < -9223372036854775808 || a-b > 9223372036854775807 || a*b > 9223372036854775807 || a*b < -9223372036854775808 {
		return
	}

	switch {
	case operator == "+":
		fmt.Println(a + b)
	case operator == "-":
		fmt.Println(a - b)
	case operator == "*":
		fmt.Println(a * b)
	case operator == "/":
		if a > b && b == 0 {
			fmt.Println("No division by 0")
		} else if a < b && a == 0 {
			fmt.Println("No division by 0")
		} else if a == 0 && b == 0 {
			fmt.Println("No division by 0")
		} else {
			fmt.Println(a / b)
		}
	case operator == "%":
		if a > b && b == 0 {
			fmt.Println("No modulo by 0")
		} else if a < b && a == 0 {
			fmt.Println("No modulo by 0")
		} else if a == 0 && b == 0 {
			fmt.Println("No modulo by 0")
		} else {
			fmt.Println(a % b)
		}
	default:
		//
	}
	fmt.Println()
}
