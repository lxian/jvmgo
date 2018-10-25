package reference

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type INVOKE_VIRTUAL struct {
	instruction.Index16Instruction
}

func (*INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	// TODO implemention
	frame.OperandStack().PopRef()
}



