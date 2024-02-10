package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func AddNodeAtEnd(head *Node, data int) *Node {
	newNode := &Node{
		data: data,
		next: nil,
	}

	if head == nil {
		return newNode
	}

	currentNode := head
	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = newNode
	return head
}

func AddNodeAtBeginning(head *Node, data int) *Node {
	newNode := &Node{
		data: data,
		next: head,
	}

	return newNode
}

func PrintLinkedList(head *Node) {
	currentNode := head
	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println()
}

func main() {
	var head *Node

	// Add nodes at the end
	head = AddNodeAtEnd(head, 10)
	head = AddNodeAtEnd(head, 20)
	head = AddNodeAtEnd(head, 30)

	fmt.Println("Linked List (after adding nodes at the end): ")
	PrintLinkedList(head)

	// Add a node at the beginning
	head = AddNodeAtBeginning(head, 99)

	fmt.Println("Linked list (after adding a node at the beginning ): ")
	PrintLinkedList(head)
}
