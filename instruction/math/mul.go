package math

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type IMUL struct {
	instruction.NoOperandsInstruction
}

func (*IMUL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(frame.OperandStack().PopInt() * frame.OperandStack().PopInt())
}

type LMUL struct {
	instruction.NoOperandsInstruction
}

func (*LMUL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(frame.OperandStack().PopLong() * frame.OperandStack().PopLong())
}

type FMUL struct {
	instruction.NoOperandsInstruction
}

func (*FMUL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(frame.OperandStack().PopFloat() * frame.OperandStack().PopFloat())
}

type DMUL struct {
	instruction.NoOperandsInstruction
}

func (*DMUL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(frame.OperandStack().PopDouble() * frame.OperandStack().PopDouble())
}
