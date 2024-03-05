package piscine 

func LastRune(s string) rune {
  for idx, char := range s{
    if idx == len(s)-1{
      return char
    }
  }
  return rune('\n')
}