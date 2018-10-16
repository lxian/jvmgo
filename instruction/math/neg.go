package math

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type INEG struct {
	instruction.NoOperandsInstruction
}
type LNEG struct {
	instruction.NoOperandsInstruction
}
type DNEG struct {
	instruction.NoOperandsInstruction
}
type FNEG struct {
	instruction.NoOperandsInstruction
}

func (*INEG) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	v = -v
	frame.OperandStack().PushInt(v)
}

func (*LNEG) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopLong()
	v = -v
	frame.OperandStack().PushLong(v)
}

func (*DNEG) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopDouble()
	v = -v
	frame.OperandStack().PushDouble(v)
}

func (*FNEG) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopFloat()
	v = -v
	frame.OperandStack().PushFloat(v)
}
