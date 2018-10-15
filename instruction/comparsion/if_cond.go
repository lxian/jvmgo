package comparsion

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
	"github.com/derekparker/delve/pkg/dwarf/frame"
)

type ifeq struct { instruction.Index16Instruction }
type ifne struct { instruction.Index16Instruction }
type iflt struct { instruction.Index16Instruction }
type ifge struct { instruction.Index16Instruction }
type ifgt struct { instruction.Index16Instruction }
type ifle struct { instruction.Index16Instruction }

func (inst *ifeq) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v == 0 {
		instruction.Branch(frame, int(inst.Index))
	}
}

func (inst *ifne) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v != 0 {
		instruction.Branch(frame, int(inst.Index))
	}
}

func (inst *iflt) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v < 0 {
		instruction.Branch(frame, int(inst.Index))
	}
}

func (inst *ifge) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v >= 0 {
		instruction.Branch(frame, int(inst.Index))
	}
}

func (inst *ifgt) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v > 0 {
		instruction.Branch(frame, int(inst.Index))
	}
}

func (inst *ifle) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	if v <= 0 {
		instruction.Branch(frame, int(inst.Index))
	}
}

