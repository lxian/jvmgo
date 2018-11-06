package array

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type AASTORE struct {
	instruction.NoOperandsInstruction
}
type BASTORE struct {
	instruction.NoOperandsInstruction
}
type CASTORE struct {
	instruction.NoOperandsInstruction
}
type DASTORE struct {
	instruction.NoOperandsInstruction
}
type FASTORE struct {
	instruction.NoOperandsInstruction
}
type IASTORE struct {
	instruction.NoOperandsInstruction
}
type LASTORE struct {
	instruction.NoOperandsInstruction
}
type SASTORE struct {
	instruction.NoOperandsInstruction
}

func (*AASTORE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopRef()
	idx := frame.OperandStack().PopInt()
	arr := frame.OperandStack().PopRef()

	arr.AssertArrIdx(idx)

	arr.Refs()[idx] = val
}

func (*BASTORE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	idx := frame.OperandStack().PopInt()
	arr := frame.OperandStack().PopRef()

	arr.AssertArrIdx(idx)

	arr.Bytes()[idx] = int8(val)
}

func (*CASTORE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	idx := frame.OperandStack().PopInt()
	arr := frame.OperandStack().PopRef()
	arr.AssertArrIdx(idx)
	arr.Chars()[idx] = int16(val)
}

func (*DASTORE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopDouble()
	idx := frame.OperandStack().PopInt()
	arr := frame.OperandStack().PopRef()
	arr.AssertArrIdx(idx)
	arr.Doubles()[idx] = val
}

func (*FASTORE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopFloat()
	idx := frame.OperandStack().PopInt()
	arr := frame.OperandStack().PopRef()
	arr.AssertArrIdx(idx)
	arr.Floats()[idx] = val
}

func (*IASTORE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	idx := frame.OperandStack().PopInt()
	arr := frame.OperandStack().PopRef()
	arr.AssertArrIdx(idx)
	arr.Ints()[idx] = val
}

func (*LASTORE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopLong()
	idx := frame.OperandStack().PopInt()
	arr := frame.OperandStack().PopRef()
	arr.AssertArrIdx(idx)
	arr.Longs()[idx] = val
}

func (*SASTORE) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	idx := frame.OperandStack().PopInt()
	arr := frame.OperandStack().PopRef()
	arr.AssertArrIdx(idx)
	arr.Shorts()[idx] = int16(val)
}
