package main

import (
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 3 {
		return
	}

	arg1 := args[0]
	arg2 := args[2]

	value1 := atoi(arg1)
	value2 := atoi(arg2)

	var result int
	switch args[1] {
	case "+":
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "*":
		result = value1 * value2
	case "/":
		if value2 == 0 {
			os.Stdout.WriteString("No division by 0")
			return
		}
		result = value1 / value2
	case "%":
		if value2 == 0 {
			os.Stdout.WriteString("No modulo by 0")
			return
		}
		result = value1 % value2
	default:
		return
	}

	// Print the result
	os.Stdout.WriteString(itoa(result))
}

func itoa(n int) (res string) {
	for n != 0 {
		res = string(rune(n%10+'0')) + res
		n = n / 10
	}
	return res
}

func atoi(s string) int {
	result := 0

	for _, char := range s {
		digit := int(char - '0')
		result = result*10 + digit
	}

	return result
}
