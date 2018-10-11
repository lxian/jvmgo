package rtda

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{}
}

func (thread *Thread) PC() int {
	return thread.pc
}

func (thread *Thread) SetPC(newPc int) {
	thread.pc = newPc
}

func (thread *Thread) PushFrame(Frame *Frame) {
	thread.stack.push(Frame)
}

func (thread *Thread) PopFrame() *Frame {
	return thread.stack.pop()
}

func (thread *Thread) CurrentFrame() *Frame {
	return thread.stack.top()
}

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{maxSize: maxSize, size: 0}
}

func (stack *Stack) push(frame *Frame) {
	if stack.size == stack.maxSize {
		panic("Thread stack overflows")
	}
	frame.lower = stack._top
	stack._top = frame
	stack.size += 1
}

func (stack *Stack) pop() *Frame {
	if stack.size < 1 {
		panic("Thread stack is emtpy")
	}
	topFrame := stack._top
	stack._top = topFrame.lower
	return topFrame
}

func (stack *Stack) top() *Frame {
	return stack._top
}
