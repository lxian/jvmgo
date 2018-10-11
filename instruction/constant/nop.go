package constant

import "jvmgo/rtda"

type NOP struct { NoOperandsInstruction }

func (inst *NOP) Execute(frame *rtda.Frame) {
}



