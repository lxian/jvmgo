package rtda

import "jvmgo/rtda/heap"

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	method       *heap.Method
	nextPC       int
}

func NewFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
		thread:       thread,
		method:       method,
	}
}

func (frame *Frame) LocalVars() LocalVars {
	return frame.localVars
}

func (frame *Frame) OperandStack() *OperandStack {
	return frame.operandStack
}

func (frame *Frame) Thread() *Thread {
	return frame.thread
}

func (frame *Frame) NextPC() int {
	return frame.nextPC
}

func (frame *Frame) SetNextPC(nextPC int) {
	frame.nextPC = nextPC
}

func (frame *Frame) Method() *heap.Method {
	return frame.method
}
