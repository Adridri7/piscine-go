package piscine

func Fibonacci(index int) int {
	if index <= 1 {
		return index
	} else if index < 0 {
		return -1
	}
	return Fibonacci(index-2) + Fibonacci(index-1)
}
