package stack

import "errors"

type Node struct {
	val  int
	next *Node
}

type LinkedStack struct {
	head *Node
	size int
}

func (stack *LinkedStack) Push(val int) {
	newNode := &Node{val: val, next: nil}
	if stack.head == nil {
		stack.head = newNode
	} else {
		newNode.next = stack.head
		stack.head = newNode
	}
	stack.size++
}

func (stack *LinkedStack) Pop() (int, error) {
	if stack.size > 0 {
		current := stack.head
		if stack.size > 1 {
			stack.head = stack.head.next
		} else {
			stack.head = nil
		}
		stack.size--
		return current.val, nil
	}
	return -1, errors.New("stack is empty")
}

func (stack *LinkedStack) Peek() (int, error) {
	if stack.size > 0 {
		return stack.head.val, nil
	}
	return -1, errors.New("stack is empty")
}

func (stack *LinkedStack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *LinkedStack) Size() int {
	return stack.size
}
