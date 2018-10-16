package comparsion

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type IFEQ struct{ instruction.BranchInstruction }
type IFNE struct{ instruction.BranchInstruction }
type IFLT struct{ instruction.BranchInstruction }
type IFGE struct{ instruction.BranchInstruction }
type IFGT struct{ instruction.BranchInstruction }
type IFLE struct{ instruction.BranchInstruction }

func (inst *IFEQ) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v == 0 {
		instruction.Branch(frame, int(inst.Offset))
	}
}

func (inst *IFNE) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v != 0 {
		instruction.Branch(frame, int(inst.Offset))
	}
}

func (inst *IFLT) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v < 0 {
		instruction.Branch(frame, int(inst.Offset))
	}
}

func (inst *IFGE) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v >= 0 {
		instruction.Branch(frame, int(inst.Offset))
	}
}

func (inst *IFGT) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v > 0 {
		instruction.Branch(frame, int(inst.Offset))
	}
}

func (inst *IFLE) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v <= 0 {
		instruction.Branch(frame, int(inst.Offset))
	}
}
