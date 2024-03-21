package piscine

type NodeI struct {
	Data int
	Next *NodeI
}

func ListSort(l *NodeI) *NodeI {
	if l == nil || l.Next == nil {
		return l
	}

	swapped := true
	for swapped {
		swapped = false
		prev := l
		current := l.Next
		for current != nil {
			if prev.Data > current.Data {
				prev.Data, current.Data = current.Data, prev.Data
				swapped = true
			}
			prev = current
			current = current.Next
		}
	}
	return l
}
