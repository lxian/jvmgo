package comparsion

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type IFNONNULL struct{ instruction.Index16Instruction }

func (inst *IFNONNULL) Execute(frame *rtda.Frame) {
	if frame.OperandStack().PopRef() != nil {
		instruction.Branch(frame, int(inst.Index))
	}
}

type IFNULL struct{ instruction.Index16Instruction }

func (inst *IFNULL) Execute(frame *rtda.Frame) {
	if frame.OperandStack().PopRef() == nil {
		instruction.Branch(frame, int(inst.Index))
	}
}
