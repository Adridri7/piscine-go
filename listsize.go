package piscine 

func ListSize(l *List) int {
  cpt := 0
  n := l.Head
  for n != nil {
    cpt ++
    n = n.Next
  }
  return cpt
}