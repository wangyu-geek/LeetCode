package stack

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

//Create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

//Len()
func (this *Stack) Len() int {
	return this.length
}

//Peek()
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

//Pop()
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}
	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

//Push
func (this *Stack) Push(value interface{}) {
	n := &node{value: value, prev: this.top}
	this.top = n
	this.length++
}
