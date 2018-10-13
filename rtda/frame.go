package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread *Thread
	nextPC int
}

func NewFrame(maxLocalVarSize uint, maxOperandStackSize uint, thread *Thread) *Frame {
	return &Frame{
		localVars: newLocalVars(maxLocalVarSize),
		operandStack: newOperandStack(maxOperandStackSize),
		thread:thread,
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

func (frame *Frame) SetNextPC(nextPC int) *Thread {
	frame.nextPC = nextPC
}
