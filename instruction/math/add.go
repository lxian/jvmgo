package math

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type IADD struct {
	instruction.NoOperandsInstruction
}

func (*IADD) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(frame.OperandStack().PopInt() + frame.OperandStack().PopInt())
}

type LADD struct {
	instruction.NoOperandsInstruction
}

func (*LADD) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(frame.OperandStack().PopLong() + frame.OperandStack().PopLong())
}

type FADD struct {
	instruction.NoOperandsInstruction
}

func (*FADD) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(frame.OperandStack().PopFloat() + frame.OperandStack().PopFloat())
}

type DADD struct {
	instruction.NoOperandsInstruction
}

func (*DADD) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(frame.OperandStack().PopDouble() + frame.OperandStack().PopDouble())
}
