package store

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

func _astore(frame *rtda.Frame, index uint) {
	value := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, value)
}

type ASTORE struct{ instruction.Index8Instruction }

func (inst *ASTORE) Execute(frame *rtda.Frame) {
	_astore(frame, inst.Index)
}

type ASTORE_0 struct {
	instruction.NoOperandsInstruction
}
type ASTORE_1 struct {
	instruction.NoOperandsInstruction
}
type ASTORE_2 struct {
	instruction.NoOperandsInstruction
}
type ASTORE_3 struct {
	instruction.NoOperandsInstruction
}

func (*ASTORE_0) Execute(frame *rtda.Frame) {
	_astore(frame, 0)
}

func (*ASTORE_1) Execute(frame *rtda.Frame) {
	_astore(frame, 1)
}

func (*ASTORE_2) Execute(frame *rtda.Frame) {
	_astore(frame, 2)
}

func (*ASTORE_3) Execute(frame *rtda.Frame) {
	_astore(frame, 3)
}
