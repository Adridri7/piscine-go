package main

import "fmt"

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

func main() {
	var link *Node

	link = insert(link, 0)
	link = insert(link, 1)
	link = insert(link, 2)
	link = insert(link, 3)
	printList(link)
}
