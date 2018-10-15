package comparsion

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type iFACMP struct {
	instruction.Index16Instruction
	shouldBranch func(v1 *rtda.Object, v2 *rtda.Object) bool
}

func (inst *iFACMP) Execute(frame *rtda.Frame) {
	v2 := frame.OperandStack().PopRef()
	v1 := frame.OperandStack().PopRef()
	if inst.shouldBranch(v1, v2) {
		instruction.Branch(frame, int(inst.Index))
	}
}

type IFACMPEQ struct { iFACMP }
type IFACMPNE struct { iFACMP }

func NewIFACMPEQ() *IFACMPEQ {
	return &IFACMPEQ{iFACMP{shouldBranch: func(v1 *rtda.Object, v2 *rtda.Object) bool {
		return v1 == v2
	}}}
}

func NewIFACMPNE() *IFACMPNE {
	return &IFACMPNE{iFACMP{shouldBranch: func(v1 *rtda.Object, v2 *rtda.Object) bool {
		return v1 != v2
	}}}
}
