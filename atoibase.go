package piscine

func AtoiBase(s string, base string) int {
  res := 0
  if !isValidBase(base) {
		return res
  }
  cpt := len(s)-1
  idx := 0
  for _, char := range s{
    for ind, c := range base {
      if char == c{
        idx = ind
      }
    }
    res += idx*RecursivePower(len(base), cpt)
    cpt--
  }
  return res
}

func RecursivePower(nb int, power int) (res int) {
	if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	} else {
		return nb * RecursivePower(nb, power-1)
	}
}

func isValidBase(base string) bool {
	if len(base) < 2 {
		return false
	}
	for _, val := range base {
		if val == '-' || val == '+' {
			return false
		} else if !contain1(base, val) {
			return false
		}
	}
	return true
}

func contain1(base string, r rune) bool {
	cpt := 0
	for _, val := range base {
		if val == r {
			cpt++
		}
	}
	return cpt == 1
}
