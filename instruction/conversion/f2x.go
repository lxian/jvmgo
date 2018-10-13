package conversion

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type F2D struct { instruction.NoOperandsInstruction }
type F2L struct { instruction.NoOperandsInstruction }
type F2I struct { instruction.NoOperandsInstruction }

func (*F2D) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopFloat()
	frame.OperandStack().PushDouble(float64(v))
}

func (*F2L) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopFloat()
	frame.OperandStack().PushLong(int64(v))
}

func (*F2I) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopFloat()
	frame.OperandStack().PushInt(int32(v))
}

