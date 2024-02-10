package main

import "fmt"

// Tạo struct Node bao gồm : data và địa chỉ Node tiếp theo
type Node struct {
	data int
	next *Node
}

func ArrayToLinkedList(arr []int) *Node {
	var head, tail *Node
	for _, val := range arr {
		newNode := &Node{
			data: val,
			next: nil,
		}

		if head == nil {
			head = newNode
			tail = newNode
		} else {
			tail.next = newNode
			tail = newNode
		}
	}

	return head
}

func PrintLinkedList(head *Node) {
	currentNode := head
	for currentNode != nil {
		fmt.Printf("%d", currentNode.data)
		currentNode = currentNode.next
	}
}

func main() {
	arr := []int{10, 20, 30, 40, 50}
	linkedList := ArrayToLinkedList(arr)

	PrintLinkedList(linkedList)
}
