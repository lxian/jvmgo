package conversion

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type I2B struct { instruction.NoOperandsInstruction }
type I2C struct { instruction.NoOperandsInstruction }
type I2F struct { instruction.NoOperandsInstruction }
type I2D struct { instruction.NoOperandsInstruction }
type I2L struct { instruction.NoOperandsInstruction }
type I2S struct { instruction.NoOperandsInstruction }

func (*I2B) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	frame.OperandStack().PushInt(int32(byte(v)))
}

func (*I2C) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	frame.OperandStack().PushInt(int32(uint16(v)))
}

func (*I2F) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	frame.OperandStack().PushFloat(float32(v))
}

func (*I2D) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	frame.OperandStack().PushDouble(float64(v))
}

func (*I2L) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	frame.OperandStack().PushLong(int64(v))
}

func (*I2S) Execute(frame *rtda.Frame) {
	v := frame.OperandStack().PopInt()
	frame.OperandStack().PushInt(int32(int16(v)))
}


