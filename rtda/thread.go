package rtda

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{stack: newStack(1024)}
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

func (thread *Thread) IsStackEmpty() bool {
	return thread.stack.top() == nil
}

func (thread *Thread) StackSize() uint {
	return thread.stack.size
}

type StackIterator struct {
	cur *Frame
}

func NewStackIterator(thread *Thread) *StackIterator {
	return &StackIterator{cur: thread.stack._top}
}

func (it *StackIterator) HasNext() bool {
	return it.cur != nil
}

func (it *StackIterator) Next() *Frame {
	top := it.cur
	it.cur = top.lower
	return top
}

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func (stack *Stack) Size() uint {
	return stack.size
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
	stack.size -= 1
	return topFrame
}

func (stack *Stack) top() *Frame {
	return stack._top
}
