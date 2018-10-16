package load

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

func _fload(frame *rtda.Frame, index uint) {
	value := frame.LocalVars().GetFloat(uint(index))
	frame.OperandStack().PushFloat(value)
}

type FLOAD struct{ instruction.Index8Instruction }

func (inst *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, inst.Index)
}

type FLOAD_0 struct {
	instruction.NoOperandsInstruction
}
type FLOAD_1 struct {
	instruction.NoOperandsInstruction
}
type FLOAD_2 struct {
	instruction.NoOperandsInstruction
}
type FLOAD_3 struct {
	instruction.NoOperandsInstruction
}

func (*FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

func (*FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

func (*FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

func (*FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}
