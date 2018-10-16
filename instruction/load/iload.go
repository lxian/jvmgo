package load

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

func _iload(frame *rtda.Frame, index uint) {
	value := frame.LocalVars().GetInt(uint(index))
	frame.OperandStack().PushInt(value)
}

type ILOAD struct{ instruction.Index8Instruction }

func (inst *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, inst.Index)
}

type ILOAD_0 struct {
	instruction.NoOperandsInstruction
}
type ILOAD_1 struct {
	instruction.NoOperandsInstruction
}
type ILOAD_2 struct {
	instruction.NoOperandsInstruction
}
type ILOAD_3 struct {
	instruction.NoOperandsInstruction
}

func (*ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

func (*ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

func (*ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

func (*ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}
