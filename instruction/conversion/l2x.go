package conversion

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type L2F struct { instruction.NoOperandsInstruction }
type L2D struct { instruction.NoOperandsInstruction }
type L2I struct { instruction.NoOperandsInstruction }

func (*L2F) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopLong()
	frame.OperandStack().PushFloat(float32(v))
}

func (*L2D) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopLong()
	frame.OperandStack().PushDouble(float64(v))
}

func (*L2I) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopLong()
	frame.OperandStack().PushInt(int32(v))
}

