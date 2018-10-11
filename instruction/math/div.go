package math

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type IDIV struct { instruction.NoOperandsInstruction }

func (*IDIV) Execute(frame *rtda.Frame) {
	v2 := frame.OperandStack().PopInt()
	v1 := frame.OperandStack().PopInt()
	if v2 == 0 {
		panic("/ by zero")
	}
	frame.OperandStack().PushInt(v1 / v2)
}

type LDIV struct { instruction.NoOperandsInstruction }

func (*LDIV) Execute(frame *rtda.Frame) {
	v2 := frame.OperandStack().PopLong()
	v1 := frame.OperandStack().PopLong()
	if v2 == 0 {
		panic("/ by zero")
	}
	frame.OperandStack().PushLong(v1 / v2)
}

type FDIV struct { instruction.NoOperandsInstruction }

func (*FDIV) Execute(frame *rtda.Frame) {
	v2 := frame.OperandStack().PopFloat()
	v1 := frame.OperandStack().PopFloat()
	if v2 == 0 {
		panic("/ by zero")
	}
	frame.OperandStack().PushFloat(v1 / v2)
}

type DDIV struct { instruction.NoOperandsInstruction }

func (*DDIV) Execute(frame *rtda.Frame) {
	v2 := frame.OperandStack().PopDouble()
	v1 := frame.OperandStack().PopDouble()
	if v2 == 0 {
		panic("/ by zero")
	}
	frame.OperandStack().PushDouble(v1 / v2)
}

