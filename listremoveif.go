package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}
	if l.Head == nil {
		l.Tail = nil
		return
	}
	current := l.Head
	for current.Next != nil {
		if current.Next.Data == data_ref {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}
	l.Tail = current
}
