package main

import "fmt"

type Node struct {
	data int
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

func (dll *DoublyLinkedList) AddNode(data int) {
	newNode := &Node{
		data: data,
		prev: nil,
		next: nil,
	}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	}
}

func (dll *DoublyLinkedList) PrintForward() {
	currentNode := dll.head
	if currentNode == nil {
		fmt.Println("Doubly Linked List is empty! ")
	}

	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.next
	}
}

func (dll *DoublyLinkedList) PrintReverse() {
	currentNode := dll.tail
	if currentNode == nil {
		fmt.Println("Doubly Linked List is empty! ")
	}

	for currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		currentNode = currentNode.prev
	}
}

func main() {
	dll := DoublyLinkedList{}
	dll.AddNode(10)
	dll.AddNode(20)
	dll.AddNode(30)
	dll.AddNode(40)

	fmt.Println("Doubly Linked List (forward): ")
	dll.PrintForward()

	fmt.Println("Doubly Linked List (Reverse): ")
	dll.PrintReverse()
}
