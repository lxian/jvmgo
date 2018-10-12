package math

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

func shiftBitsCnt(frame *rtda.Frame) uint {
	return uint(frame.OperandStack().PopInt() & 0x0000001F) // take low 5 bits only
}

type ishl struct { instruction.NoOperandsInstruction }
type ishr struct { instruction.NoOperandsInstruction }
type iushr struct { instruction.NoOperandsInstruction }

func (*ishl) Execute(frame *rtda.Frame) {
	v2 := shiftBitsCnt(frame)
	v1 := frame.OperandStack().PopInt()
	shifted := v1 << v2
	frame.OperandStack().PushInt(shifted)
}

func (*ishr) Execute(frame *rtda.Frame) {
	v2 := shiftBitsCnt(frame)
	v1 := frame.OperandStack().PopInt()

	shifted := v1 >> v2

	if (v1 & 0x80000000 > 0) { // first bit is 1
		var extendedSignBits int32 = 0xffffffff << (32 - v2)
		shifted = shifted | extendedSignBits
	}

	frame.OperandStack().PushInt(shifted)
}

func (*iushr) Execute(frame *rtda.Frame) {
	v2 := shiftBitsCnt(frame)
	v1 := frame.OperandStack().PopInt()
	shifted := v1 >> v2
	frame.OperandStack().PushInt(shifted)
}

type lshl struct { instruction.NoOperandsInstruction }
type lshr struct { instruction.NoOperandsInstruction }
type lushr struct { instruction.NoOperandsInstruction }

func (*lshl) Execute(frame *rtda.Frame) {
	v2 := shiftBitsCnt(frame)
	v1 := frame.OperandStack().PopLong()
	shifted := v1 << v2
	frame.OperandStack().PushLong(shifted)
}

func (*lshr) Execute(frame *rtda.Frame) {
	v2 := shiftBitsCnt(frame)
	v1 := frame.OperandStack().PopLong()

	shifted := v1 >> v2

	if (v1 & 0x80000000 > 0) { // first bit is 1
		var extendedSignBits int64 = 0xffffffffffffffff << (64 - v2)
		shifted = shifted | extendedSignBits
	}

	frame.OperandStack().PushLong(shifted)
}

func (*lushr) Execute(frame *rtda.Frame) {
	v2 := shiftBitsCnt(frame)
	v1 := frame.OperandStack().PopLong()
	shifted := v1 >> v2
	frame.OperandStack().PushLong(shifted)
}
