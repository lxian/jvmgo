package math

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type IAND struct { instruction.NoOperandsInstruction }
type IOR struct { instruction.NoOperandsInstruction }
type IXOR struct { instruction.NoOperandsInstruction }

func (*IAND) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(frame.OperandStack().PopInt() & frame.OperandStack().PopInt())
}

func (*IOR) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(frame.OperandStack().PopInt() | frame.OperandStack().PopInt())
}

func (*IXOR) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(frame.OperandStack().PopInt() ^ frame.OperandStack().PopInt())
}

type LAND struct { instruction.NoOperandsInstruction }
type LOR struct { instruction.NoOperandsInstruction }
type LXOR struct { instruction.NoOperandsInstruction }

func (*LAND) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(frame.OperandStack().PopLong() & frame.OperandStack().PopLong())
}

func (*LOR) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(frame.OperandStack().PopLong() | frame.OperandStack().PopLong())
}

func (*LXOR) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(frame.OperandStack().PopLong() ^ frame.OperandStack().PopLong())
}
