package main

import "fmt"

type Node struct {
	data int
	next *Node
}

func DeleteNode(head *Node, value int) *Node {
	// check xem node head có giữ giá trị cần xóa không
	if head != nil && head.data == value {
		return head.next
	}

	currentNode := head
	var prevNode *Node

	for currentNode != nil && currentNode.data != value {
		prevNode = currentNode
		currentNode = currentNode.next
	}

	if currentNode != nil {
		prevNode.next = currentNode.next
	}

	return head
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

	// create a linked list: 10 -> 20 -> 30 -> 40 -> 50

	head = &Node{data: 10}
	head.next = &Node{data: 20}
	head.next.next = &Node{data: 30}
	head.next.next.next = &Node{data: 40}
	head.next.next.next.next = &Node{data: 50}

	fmt.Println("Linked list (before deletion): ")
	PrintLinkedList(head)

	// Delete a node with value 30
	head = DeleteNode(head, 30)

	fmt.Println("Linked List (after deleteion): ")
	PrintLinkedList(head)
}
