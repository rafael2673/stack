package stack

import "errors"

type ArrayStack struct {
	values []interface{}
	size   int
}

func (stack *ArrayStack) Init(size int) error {
	if size > 0 {
		stack.values = make([]interface{}, size)
		return nil
	} else {
		return errors.New("can't init ArrayStack with size <= 0")
	}
}

func (stack *ArrayStack) doubleArray() {
	curSize := len(stack.values)
	doubledValues := make([]interface{}, 2*curSize)

	for i := 0; i < curSize; i++ {
		doubledValues[i] = stack.values[i]
	}
	stack.values = doubledValues
}

func (stack *ArrayStack) Push(val interface{}) {
	if len(stack.values) == stack.size {
		stack.doubleArray()
	}
	stack.values[stack.size] = val
	stack.size++
}
func (stack *ArrayStack) Pop() (interface{}, error) {
	if stack.size > 0 {
		stack.size--
		return stack.values[stack.size], nil
	}
	return -1, errors.New("empty stack")

}
func (stack *ArrayStack) Peek() (interface{}, error) {
	if stack.size > 0 {
		return stack.values[stack.size-1], nil
	}
	return -1, errors.New("empty stack")
}
func (stack *ArrayStack) IsEmpty() bool {
	return stack.size == 0
}
func (stack *ArrayStack) Size() int {
	return stack.size
}
func Balparenteses(par string) bool {
	stack := &LinkedStack{}

	for i := 0; i < len(par); i++ {
		char := string(par[i])
		switch char {
		case "(":
			stack.Push(char)
		case ")":
			top, err := stack.Pop()
			if err != nil || top != "(" {
				return false
			}
		}
	}
	return stack.IsEmpty()
}
