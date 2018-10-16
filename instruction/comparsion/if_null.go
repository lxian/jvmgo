package comparsion

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type IFNONNULL struct{ instruction.BranchInstruction }

func (inst *IFNONNULL) Execute(frame *rtda.Frame) {
	if frame.OperandStack().PopRef() != nil {
		instruction.Branch(frame, int(inst.Offset))
	}
}

type IFNULL struct{ instruction.BranchInstruction }

func (inst *IFNULL) Execute(frame *rtda.Frame) {
	if frame.OperandStack().PopRef() == nil {
		instruction.Branch(frame, int(inst.Offset))
	}
}
