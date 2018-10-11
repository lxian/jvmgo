package constant

import "jvmgo/rtda"

// ACONST
type ACONST_NULL struct { NoOperandsInstruction }

func (*ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

// ICONST
type ICONST_M1 struct { NoOperandsInstruction }
type ICONST_0 struct { NoOperandsInstruction }
type ICONST_1 struct { NoOperandsInstruction }
type ICONST_2 struct { NoOperandsInstruction }
type ICONST_3 struct { NoOperandsInstruction }
type ICONST_4 struct { NoOperandsInstruction }
type ICONST_5 struct { NoOperandsInstruction }

func (*ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

func (*ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

func (*ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

func (*ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

func (*ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

func (*ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}

func (*ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

// FCONST
type FCONST_0 struct { NoOperandsInstruction }
type FCONST_1 struct { NoOperandsInstruction }
type FCONST_2 struct { NoOperandsInstruction }

func (*FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0)
}

func (*FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1)
}

func (*FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2)
}

// DCONST
type DCONST_0 struct { NoOperandsInstruction }
type DCONST_1 struct { NoOperandsInstruction }

func (*DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0)
}

func (*DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1)
}

// LCONST
type LCONST_0 struct { NoOperandsInstruction }
type LCONST_1 struct { NoOperandsInstruction }

func (*LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}

func (*LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}

