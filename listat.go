package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListAt(l *NodeL, pos int) *NodeL {
	if l == nil || pos < 0 {
		return nil
	}
	current := l
	idx := 0
	for current != nil {
		if idx == pos {
			return current
		}
		current = current.Next
		idx++
	}
	return nil
}
