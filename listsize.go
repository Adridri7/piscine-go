package piscine 

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListSize(l *List) int {
  cpt := 0
  n := l.Head
  for n != nil {
    cpt ++
    n = n.Next
  }
  return cpt
}