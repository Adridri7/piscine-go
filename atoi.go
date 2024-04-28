package piscine

func Atoi(s string) (res int) {
	res = 0
	sign := 1

	for idx, char := range s {
		if char >= '0' && char <= '9' || char == '-' || char == '+' {
			if idx == 0 && char == '-' {
				sign = -1
			} else if idx == 0 && char == '+' {
				sign = 1
			} else if char == '-' || char == '+' && idx > 0 {
				return 0
			} else {
				digit := char - '0'
				res = res*10 + int(digit)
			}
		} else {
			return 0
		}
	}
	return res * sign
}
