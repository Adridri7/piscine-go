package piscine 

type nodeL struct {
	Data interface{}
	Next *nodeL
}

type list struct {
	Head *nodeL
	Tail *nodeL
}

func ListSize(l *list) int {
  cpt := 0
  n := l.Head
  for n != nil {
    n = n.Next
    cpt ++
  }
  return cpt
}