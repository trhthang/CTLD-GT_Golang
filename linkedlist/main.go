package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func printLinkedList(head *Node) {
	if head == nil {
		fmt.Println("LinkedList is empty")
	} else {
		temp := head
		for temp != nil {
			fmt.Print(temp.value, " ")
			temp = temp.next
		}
		fmt.Println()
	}
}

// Thêm phần tử vào đầu của LinkedList
func addToHead(head *Node, value int) *Node {
	newNode := &Node{value: value}

	if head != nil {
		newNode.next = head
	}

	return newNode
}

// Thêm phần từ vào cuối của LinkedList
func addToTail(head *Node, value int) *Node {
	newNode := &Node{value: value}
	if head == nil {
		return newNode
	}
	lastNode := head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}

	lastNode.next = newNode

	return head
}

// Thêm 1 phần tử ở giữa
func addToIndex(head *Node, value int, index int) *Node {
	if index == 0 {
		addToHead(head, value)
	}

	newNode := &Node{value: value}
	curNode := head
	count := 0
	for curNode != nil {
		count++
		if count == index {
			newNode.next = curNode.next
			curNode.next = newNode
			break

		}
		curNode = curNode.next
	}

	return head
}

// xoa phan tu o dau
func removeAthead(head *Node) *Node {
	if head != nil {
		return head.next
	}

	return head
}

// xóa phần tử ở cuối
func removeAtTail(head *Node) *Node {
	if head == nil {
		return nil
	}

	lastNode := head
	var prevNode *Node

	for lastNode.next != nil {
		prevNode = lastNode
		lastNode = lastNode.next
	}

	if prevNode == nil {
		return nil
	} else {
		prevNode.next = nil
	}

	return head
}

// Xoa phan tu o giua
func removeAtIndex(head *Node, index int) *Node {
	if head == nil || index < 0 {
		return nil
	}

	if index == 0 {
		return removeAthead(head)
	}

	curNode := head
	var prevNode *Node
	count := 0

	for curNode != nil {
		if count == index {
			if prevNode != nil {
				prevNode.next = curNode.next
			}
			break
		}
		prevNode = curNode
		curNode = curNode.next
		count++
	}

	return head
}

func main() {

	n1 := &Node{value: 1}
	n2 := &Node{value: 2}
	n3 := &Node{value: 3}
	n4 := &Node{value: 4}

	n1.next = n2
	n2.next = n3
	n3.next = n4

	printLinkedList(n1)

	n1 = addToHead(n1, 0)

	printLinkedList(n1)

}
