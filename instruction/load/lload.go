package load

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

func _lload(frame *rtda.Frame, index uint) {
	value := frame.LocalVars().GetLong(uint(index))
	frame.OperandStack().PushLong(value)
}

type LLOAD struct{ instruction.Index8Instruction }

func (inst *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, inst.Index)
}

type LLOAD_0 struct {
	instruction.NoOperandsInstruction
}
type LLOAD_1 struct {
	instruction.NoOperandsInstruction
}
type LLOAD_2 struct {
	instruction.NoOperandsInstruction
}
type LLOAD_3 struct {
	instruction.NoOperandsInstruction
}

func (*LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}

func (*LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

func (*LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

func (*LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}
