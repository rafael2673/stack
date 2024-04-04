package stack

type IStack interface {
	Push(value interface{})
	Pop() (interface{}, error)
	Peek() (interface{}, error)
	IsEmpty() bool
	Size() int
}
