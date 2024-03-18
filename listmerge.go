package piscine

func ListMerge(l1 *List, l2 *List) {
	l1.Tail = l2.Head
}
