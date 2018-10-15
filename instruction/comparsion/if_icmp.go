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

type IFICMPEQ struct { iFICMP }
type IFICMPNE struct { iFICMP }
type IFICMPLT struct { iFICMP }
type IFICMPGE struct { iFICMP }
type IFICMPGT struct { iFICMP }
type IFICMPLE struct { iFICMP }

func NewIFICMPEQ() *IFICMPEQ {
	return &IFICMPEQ{iFICMP{shouldBranch:func(v1 int32, v2 int32) bool { return v1 == v2 }}}
}

func NewIFICMPNE() *IFICMPNE {
	return &IFICMPNE{iFICMP{shouldBranch:func(v1 int32, v2 int32) bool { return v1 != v2 }}}
}


func NewIFICMPLT() *IFICMPLT {
	return &IFICMPLT{iFICMP{shouldBranch:func(v1 int32, v2 int32) bool { return v1 < v2 }}}
}


func NewIFICMPGE() *IFICMPGE {
	return &IFICMPGE{iFICMP{shouldBranch:func(v1 int32, v2 int32) bool { return v1 >= v2 }}}
}

func NewIFICMPGT() *IFICMPGT {
	return &IFICMPGT{iFICMP{shouldBranch:func(v1 int32, v2 int32) bool { return v1 > v2 }}}
}

func NewIFICMPLE() *IFICMPLE {
	return &IFICMPLE{iFICMP{shouldBranch:func(v1 int32, v2 int32) bool { return v1 <= v2 }}}
}


