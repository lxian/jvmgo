package math

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type ISUB struct { instruction.NoOperandsInstruction }

func (*ISUB) Execute(frame *rtda.Frame) {
	v2 := frame.OperandStack().PopInt()
	v1 := frame.OperandStack().PopInt()
	frame.OperandStack().PushInt(v1 - v2)
}

type LSUB struct { instruction.NoOperandsInstruction }

func (*LSUB) Execute(frame *rtda.Frame) {
	v2 := frame.OperandStack().PopLong()
	v1 := frame.OperandStack().PopLong()
	frame.OperandStack().PushLong(v1 - v2)
}

type FSUB struct { instruction.NoOperandsInstruction }

func (*FSUB) Execute(frame *rtda.Frame) {
	v2 := frame.OperandStack().PopFloat()
	v1 := frame.OperandStack().PopFloat()
	frame.OperandStack().PushFloat(v1 - v2)
}

type DSUB struct { instruction.NoOperandsInstruction }

func (*DSUB) Execute(frame *rtda.Frame) {
	v2 := frame.OperandStack().PopDouble()
	v1 := frame.OperandStack().PopDouble()
	frame.OperandStack().PushDouble(v1 - v2)
}

