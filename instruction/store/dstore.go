package store

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

func _dstore(frame *rtda.Frame, index uint) {
	value := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, value)
}

type DSTORE struct{ instruction.Index8Instruction }

func (inst *DSTORE) Execute(frame *rtda.Frame) {
	_dstore(frame, inst.Index)
}

type DSTORE_0 struct {
	instruction.NoOperandsInstruction
}
type DSTORE_1 struct {
	instruction.NoOperandsInstruction
}
type DSTORE_2 struct {
	instruction.NoOperandsInstruction
}
type DSTORE_3 struct {
	instruction.NoOperandsInstruction
}

func (*DSTORE_0) Execute(frame *rtda.Frame) {
	_dstore(frame, 0)
}

func (*DSTORE_1) Execute(frame *rtda.Frame) {
	_dstore(frame, 1)
}

func (*DSTORE_2) Execute(frame *rtda.Frame) {
	_dstore(frame, 2)
}

func (*DSTORE_3) Execute(frame *rtda.Frame) {
	_dstore(frame, 3)
}
