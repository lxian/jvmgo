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
	frame.OperandStack().PushInt(v1 >> v2)
}

func (*iushr) Execute(frame *rtda.Frame) {
	v2 := shiftBitsCnt(frame)
	v1 := frame.OperandStack().PopInt()
	shifted := uint32(v1) >> v2
	frame.OperandStack().PushInt(int32(shifted))
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
	frame.OperandStack().PushLong(v1 >> v2)
}

func (*lushr) Execute(frame *rtda.Frame) {
	v2 := shiftBitsCnt(frame)
	v1 := frame.OperandStack().PopLong()
	shifted := uint64(v1) >> v2
	frame.OperandStack().PushLong(int64(shifted))
}
