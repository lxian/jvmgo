package comparsion

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type iFACMP struct {
	instruction.BranchInstruction
	shouldBranch func(v1 *rtda.Object, v2 *rtda.Object) bool
}

func (inst *iFACMP) Execute(frame *rtda.Frame) {
	v2 := frame.OperandStack().PopRef()
	v1 := frame.OperandStack().PopRef()
	if inst.shouldBranch(v1, v2) {
		instruction.Branch(frame, int(inst.Offset))
	}
}

type IF_ACMPEQ struct{ iFACMP }
type IF_ACMPNE struct{ iFACMP }

func NewIFACMPEQ() *IF_ACMPEQ {
	return &IF_ACMPEQ{iFACMP{shouldBranch: func(v1 *rtda.Object, v2 *rtda.Object) bool {
		return v1 == v2
	}}}
}

func NewIFACMPNE() *IF_ACMPNE {
	return &IF_ACMPNE{iFACMP{shouldBranch: func(v1 *rtda.Object, v2 *rtda.Object) bool {
		return v1 != v2
	}}}
}
