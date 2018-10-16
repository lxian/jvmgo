package store

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

func _fstore(frame *rtda.Frame, index uint) {
	value := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, value)
}

type FSTORE struct{ instruction.Index8Instruction }

func (inst *FSTORE) Execute(frame *rtda.Frame) {
	_fstore(frame, inst.Index)
}

type FSTORE_0 struct {
	instruction.NoOperandsInstruction
}
type FSTORE_1 struct {
	instruction.NoOperandsInstruction
}
type FSTORE_2 struct {
	instruction.NoOperandsInstruction
}
type FSTORE_3 struct {
	instruction.NoOperandsInstruction
}

func (*FSTORE_0) Execute(frame *rtda.Frame) {
	_fstore(frame, 0)
}

func (*FSTORE_1) Execute(frame *rtda.Frame) {
	_fstore(frame, 1)
}

func (*FSTORE_2) Execute(frame *rtda.Frame) {
	_fstore(frame, 2)
}

func (*FSTORE_3) Execute(frame *rtda.Frame) {
	_fstore(frame, 3)
}
