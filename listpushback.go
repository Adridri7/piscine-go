package piscine

func ListPushBack(l *List, data interface{}) {
	node := &NodeL{Data: data}

	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
}
