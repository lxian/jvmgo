package store

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

func _lstore(frame *rtda.Frame, index uint) {
	value := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, value)
}

type LSTORE struct { instruction.Index8Instruction }

func (inst *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, inst.Index)
}

type LSTORE_0 struct { instruction.NoOperandsInstruction }
type LSTORE_1 struct { instruction.NoOperandsInstruction }
type LSTORE_2 struct { instruction.NoOperandsInstruction }
type LSTORE_3 struct { instruction.NoOperandsInstruction }

func (*LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}

func (*LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}

func (*LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}

func (*LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}



