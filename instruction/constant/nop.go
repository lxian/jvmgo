package constant

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type NOP struct {
	instruction.NoOperandsInstruction
}

func (inst *NOP) Execute(frame *rtda.Frame) {
}
