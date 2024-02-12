package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func PrintLinkedList(head *Node) {
	fmt.Println("Linked List (Original Order): ")
	PrintAllElementByOriginalOrder(head)
	fmt.Println("Linked List (Reverse Order): ")
	PrintAllElementByReverseOrder(head)
	fmt.Println()
}

func PrintAllElementByOriginalOrder(head *Node) {
	// Dieu kien dung
	if head == nil {
		return
	}

	fmt.Printf("%d ", head.data)
	PrintAllElementByOriginalOrder(head.next)
}

func PrintAllElementByReverseOrder(head *Node) {
	// Dieu kien dung
	if head == nil {
		return
	}

	PrintAllElementByReverseOrder(head.next)
	fmt.Printf("%d ", head.data)
}

func main() {
	var head *Node
	head = &Node{data: 10}
	head.next = &Node{data: 20}
	head.next.next = &Node{data: 30}
	head.next.next.next = &Node{data: 40}
	head.next.next.next.next = &Node{data: 50}

	PrintLinkedList(head)
}
