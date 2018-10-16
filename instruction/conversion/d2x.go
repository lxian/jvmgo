package conversion

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type D2F struct {
	instruction.NoOperandsInstruction
}
type D2L struct {
	instruction.NoOperandsInstruction
}
type D2I struct {
	instruction.NoOperandsInstruction
}

func (*D2F) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopDouble()
	frame.OperandStack().PushFloat(float32(v))
}

func (*D2L) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopDouble()
	frame.OperandStack().PushLong(int64(v))
}

func (*D2I) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopDouble()
	frame.OperandStack().PushInt(int32(v))
}
