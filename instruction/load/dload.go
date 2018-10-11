package load

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

func _dload(frame *rtda.Frame, index uint) {
	value := frame.LocalVars().GetDouble(uint(index))
	frame.OperandStack().PushDouble(value)
}

type DLOAD struct { instruction.Index8Instruction }

func (inst *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, inst.Index)
}

type DLOAD_0 struct { instruction.NoOperandsInstruction }
type DLOAD_1 struct { instruction.NoOperandsInstruction }
type DLOAD_2 struct { instruction.NoOperandsInstruction }
type DLOAD_3 struct { instruction.NoOperandsInstruction }

func (*DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

func (*DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

func (*DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

func (*DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}



