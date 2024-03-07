package piscine

func AtoiBase(s string, base string) int {
  res := 0
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
