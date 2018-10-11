package store

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

func _istore(frame *rtda.Frame, index uint) {
	value := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, value)
}

type ISTORE struct { instruction.Index8Instruction }

func (inst *ISTORE) Execute(frame *rtda.Frame) {
	_istore(frame, inst.Index)
}

type ISTORE_0 struct { instruction.NoOperandsInstruction }
type ISTORE_1 struct { instruction.NoOperandsInstruction }
type ISTORE_2 struct { instruction.NoOperandsInstruction }
type ISTORE_3 struct { instruction.NoOperandsInstruction }

func (*ISTORE_0) Execute(frame *rtda.Frame) {
	_istore(frame, 0)
}

func (*ISTORE_1) Execute(frame *rtda.Frame) {
	_istore(frame, 1)
}

func (*ISTORE_2) Execute(frame *rtda.Frame) {
	_istore(frame, 2)
}

func (*ISTORE_3) Execute(frame *rtda.Frame) {
	_istore(frame, 3)
}



