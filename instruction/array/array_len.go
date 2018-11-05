package array

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type ARRAY_LENGTH struct {
	instruction.NoOperandsInstruction
}

func (*ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	arr := frame.OperandStack().PopRef()
	frame.OperandStack().PushInt(arr.ArrayLength())
}



