// Cách 1: Sử dụng phương pháp LẶP để in tất cả phần từ trong danh sách liên kết theo thứ tự ban đầu và thứ tự ngược lại
package main

import (
	"fmt"
	"strings"
)

type Node struct {
	data int
	next *Node
}

func PrintLinkedList(head *Node) {
	currentNode := head
	var elements []string

	for currentNode != nil {
		value := fmt.Sprintf("%d", currentNode.data)
		elements = append(elements, value)
		currentNode = currentNode.next
	}

	fmt.Println("Linked List (Original Order): ")
	fmt.Println(strings.Join(elements, " "))

	fmt.Println("Linked List (Reverse Order): ")
	for i := len(elements) - 1; i >= 0; i-- {
		fmt.Printf("%s ", elements[i])
	}
	fmt.Println()

}

func main() {
	var head *Node

	// Create a Linked List: 10  -> 20 -> 30 -> 40 -> 50
	head = &Node{data: 10}
	head.next = &Node{data: 20}
	head.next.next = &Node{data: 30}
	head.next.next.next = &Node{data: 40}
	head.next.next.next.next = &Node{data: 50}

	PrintLinkedList(head)
}
