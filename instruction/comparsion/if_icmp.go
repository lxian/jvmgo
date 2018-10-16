package comparsion

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type iFICMP struct {
	instruction.Index16Instruction
	shouldBranch func(v1 int32, v2 int32) bool
}

func (inst *iFICMP) Execute(frame *rtda.Frame) {
	v2 := frame.OperandStack().PopInt()
	v1 := frame.OperandStack().PopInt()
	if inst.shouldBranch(v1, v2) {
		instruction.Branch(frame, int(inst.Index))
	}
}

type IF_ICMPEQ struct{ iFICMP }
type IF_ICMPNE struct{ iFICMP }
type IF_ICMPLT struct{ iFICMP }
type IF_ICMPGE struct{ iFICMP }
type IF_ICMPGT struct{ iFICMP }
type IF_ICMPLE struct{ iFICMP }

func NewIFICMPEQ() *IF_ICMPEQ {
	return &IF_ICMPEQ{iFICMP{shouldBranch: func(v1 int32, v2 int32) bool { return v1 == v2 }}}
}

func NewIFICMPNE() *IF_ICMPNE {
	return &IF_ICMPNE{iFICMP{shouldBranch: func(v1 int32, v2 int32) bool { return v1 != v2 }}}
}

func NewIFICMPLT() *IF_ICMPLT {
	return &IF_ICMPLT{iFICMP{shouldBranch: func(v1 int32, v2 int32) bool { return v1 < v2 }}}
}

func NewIFICMPGE() *IF_ICMPGE {
	return &IF_ICMPGE{iFICMP{shouldBranch: func(v1 int32, v2 int32) bool { return v1 >= v2 }}}
}

func NewIFICMPGT() *IF_ICMPGT {
	return &IF_ICMPGT{iFICMP{shouldBranch: func(v1 int32, v2 int32) bool { return v1 > v2 }}}
}

func NewIFICMPLE() *IF_ICMPLE {
	return &IF_ICMPLE{iFICMP{shouldBranch: func(v1 int32, v2 int32) bool { return v1 <= v2 }}}
}
