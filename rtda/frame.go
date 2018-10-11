package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
}

func NewFrame(maxLocalVarSize uint, maxOperandStackSize uint) *Frame {
	return &Frame{localVars: newLocalVars(maxLocalVarSize), operandStack: newOperandStack(maxOperandStackSize), }
}

func (frame *Frame) LocalVars() LocalVars {
	return frame.localVars
}

func (frame *Frame) OperandStack() *OperandStack {
	return frame.operandStack
}

