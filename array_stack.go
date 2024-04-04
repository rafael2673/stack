package stack

import "errors"

type ArrayStack struct {
	values []int
	size   int
}

func (stack *ArrayStack) Init(size int) error {
	if size > 0 {
		stack.values = make([]int, size)
		return nil
	} else {
		return errors.New("Can't init ArrayStack with size <= 0")
	}
}

func (stack *ArrayStack) doubleArray() {
	curSize := len(stack.values)
	doubledValues := make([]int, 2*curSize)

	for i := 0; i < curSize; i++ {
		doubledValues[i] = stack.values[i]
	}
	stack.values = doubledValues
}

func (stack *ArrayStack) Push(val int)
func (stack *ArrayStack) Pop() (int, error)
func (stack *ArrayStack) Peek() (int, error)
func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}
func (stack *ArrayStack) Size() int {
	return stack.size
}
