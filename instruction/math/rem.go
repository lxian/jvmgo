package math

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
	"math"
)

type IREM struct {
	instruction.NoOperandsInstruction
}

func (*IREM) Execute(frame *rtda.Frame) {
	v2 := frame.OperandStack().PopInt()
	v1 := frame.OperandStack().PopInt()
	if v2 == 0 {
		panic("/ by zero")
	}
	frame.OperandStack().PushInt(v1 % v2)
}

type LREM struct {
	instruction.NoOperandsInstruction
}

func (*LREM) Execute(frame *rtda.Frame) {
	v2 := frame.OperandStack().PopLong()
	v1 := frame.OperandStack().PopLong()
	if v2 == 0 {
		panic("/ by zero")
	}
	frame.OperandStack().PushLong(v1 % v2)
}

type FREM struct {
	instruction.NoOperandsInstruction
}

func (*FREM) Execute(frame *rtda.Frame) {
	v2 := frame.OperandStack().PopFloat()
	v1 := frame.OperandStack().PopFloat()
	frame.OperandStack().PushFloat(float32(math.Mod(float64(v1), float64(v2))))
}

type DREM struct {
	instruction.NoOperandsInstruction
}

func (*DREM) Execute(frame *rtda.Frame) {
	v2 := frame.OperandStack().PopDouble()
	v1 := frame.OperandStack().PopDouble()
	frame.OperandStack().PushDouble(math.Mod(v1, v2))
}
