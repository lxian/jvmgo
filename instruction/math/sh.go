package math

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type ishl struct { instruction.NoOperandsInstruction }

func (*ishl) Execute(frame *rtda.Frame) {
	v2 := uint(frame.OperandStack().PopInt() & 0x0000001F) // take low 5 bits only
	v1 := frame.OperandStack().PopInt()
	shifted := v1 << v2
	frame.OperandStack().PushInt(shifted)
}

type ishr struct { instruction.NoOperandsInstruction }

func (*ishr) Execute(frame *rtda.Frame) {
	v2 := uint(frame.OperandStack().PopInt() & 0x0000001F) // take low 5 bits only
	v1 := frame.OperandStack().PopInt()
	shifted := v1 >> v2
	frame.OperandStack().PushInt(shifted)
}

type iushr struct { instruction.NoOperandsInstruction }

type lshl struct { instruction.NoOperandsInstruction }
type lshr struct { instruction.NoOperandsInstruction }
type lushr struct { instruction.NoOperandsInstruction }
