package main

import "fmt"

// Stack : LIFO - Last In First Out

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// Remove and return top element of stack. Return false if stack is empty
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func main() {
	var stack Stack // create a stack variable of type Stack

	stack.Push("1")
	stack.Push("2")
	stack.Push("3")
	stack.Push("4")

	stack.Pop()

	fmt.Println(stack)
}
