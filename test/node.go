package main

import (
	"fmt"
	"piscine"
)

type Node struct {
	data int
	next *Node
}

func insert(head *Node, data int) *Node {
	n := &Node{data: data}
	if head == nil {
		return n
	} else {
		n.next = head
		return n
	}
}

func printList(head *Node) {
	for head != nil {
		fmt.Print(head.data, " -> ")
		head = head.next
	}
	fmt.Println(nil)
}

/*func main() {
	link := &piscine.List{}

	piscine.ListPushFront(link, "Hello")
	piscine.ListPushFront(link, "man")
	piscine.ListPushFront(link, "how are you")

	it := link.Head
	for it != nil {
		fmt.Print(it.Data, " ")
		it = it.Next
	}
	fmt.Println()
}*/

func main() {
	root := &piscine.TreeNode{Data: "4"}
	piscine.BTreeInsertData(root, "1")
	piscine.BTreeInsertData(root, "7")
	piscine.BTreeInsertData(root, "5")
	piscine.BTreeApplyInorder(root, fmt.Println)

}
