package piscine

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
