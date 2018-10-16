package math

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

func shiftBitsCnt(frame *rtda.Frame) uint {
	return uint(frame.OperandStack().PopInt() & 0x0000001F) // take low 5 bits only
}

type ISHL struct {
	instruction.NoOperandsInstruction
}
type ISHR struct {
	instruction.NoOperandsInstruction
}
type IUSHR struct {
	instruction.NoOperandsInstruction
}

func (*ISHL) Execute(frame *rtda.Frame) {
	v2 := shiftBitsCnt(frame)
	v1 := frame.OperandStack().PopInt()
	shifted := v1 << v2
	frame.OperandStack().PushInt(shifted)
}

func (*ISHR) Execute(frame *rtda.Frame) {
	v2 := shiftBitsCnt(frame)
	v1 := frame.OperandStack().PopInt()
	frame.OperandStack().PushInt(v1 >> v2)
}

func (*IUSHR) Execute(frame *rtda.Frame) {
	v2 := shiftBitsCnt(frame)
	v1 := frame.OperandStack().PopInt()
	shifted := uint32(v1) >> v2
	frame.OperandStack().PushInt(int32(shifted))
}

type LSHL struct {
	instruction.NoOperandsInstruction
}
type LSHR struct {
	instruction.NoOperandsInstruction
}
type LUSHR struct {
	instruction.NoOperandsInstruction
}

func (*LSHL) Execute(frame *rtda.Frame) {
	v2 := shiftBitsCnt(frame)
	v1 := frame.OperandStack().PopLong()
	shifted := v1 << v2
	frame.OperandStack().PushLong(shifted)
}

func (*LSHR) Execute(frame *rtda.Frame) {
	v2 := shiftBitsCnt(frame)
	v1 := frame.OperandStack().PopLong()
	frame.OperandStack().PushLong(v1 >> v2)
}

func (*LUSHR) Execute(frame *rtda.Frame) {
	v2 := shiftBitsCnt(frame)
	v1 := frame.OperandStack().PopLong()
	shifted := uint64(v1) >> v2
	frame.OperandStack().PushLong(int64(shifted))
}
