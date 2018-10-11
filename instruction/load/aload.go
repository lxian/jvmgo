package load

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

func _aload(frame *rtda.Frame, index uint) {
	value := frame.LocalVars().GetRef(uint(index))
	frame.OperandStack().PushRef(value)
}

type ALOAD struct { instruction.Index8Instruction }

func (inst *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, inst.Index)
}

type ALOAD_0 struct { instruction.NoOperandsInstruction }
type ALOAD_1 struct { instruction.NoOperandsInstruction }
type ALOAD_2 struct { instruction.NoOperandsInstruction }
type ALOAD_3 struct { instruction.NoOperandsInstruction }

func (*ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

func (*ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

func (*ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

func (*ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}



