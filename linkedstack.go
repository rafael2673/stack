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
}

func (stack *LinkedStack) Pop() (int, error) {
	return -1, errors.New("error msg")
}

func (stack *LinkedStack) Peek() (int, error) {
	return -1, errors.New("error msg")
}

func (stack *LinkedStack) IsEmpty() bool {
	return stack.size == 0
}

func (stack *LinkedStack) Size() int {
	return stack.size
}
