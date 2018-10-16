package comparsion

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type IFEQ struct{ instruction.Index16Instruction }
type IFNE struct{ instruction.Index16Instruction }
type IFLT struct{ instruction.Index16Instruction }
type IFGE struct{ instruction.Index16Instruction }
type IFGT struct{ instruction.Index16Instruction }
type IFLE struct{ instruction.Index16Instruction }

func (inst *IFEQ) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v == 0 {
		instruction.Branch(frame, int(inst.Index))
	}
}

func (inst *IFNE) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v != 0 {
		instruction.Branch(frame, int(inst.Index))
	}
}

func (inst *IFLT) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v < 0 {
		instruction.Branch(frame, int(inst.Index))
	}
}

func (inst *IFGE) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v >= 0 {
		instruction.Branch(frame, int(inst.Index))
	}
}

func (inst *IFGT) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v > 0 {
		instruction.Branch(frame, int(inst.Index))
	}
}

func (inst *IFLE) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v <= 0 {
		instruction.Branch(frame, int(inst.Index))
	}
}
