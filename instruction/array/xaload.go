package array

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

type AALOAD struct {
	instruction.NoOperandsInstruction
}
type BALOAD struct {
	instruction.NoOperandsInstruction
}
type CALOAD struct {
	instruction.NoOperandsInstruction
}
type DALOAD struct {
	instruction.NoOperandsInstruction
}
type FALOAD struct {
	instruction.NoOperandsInstruction
}
type IALOAD struct {
	instruction.NoOperandsInstruction
}
type LALOAD struct {
	instruction.NoOperandsInstruction
}
type SALOAD struct {
	instruction.NoOperandsInstruction
}

func _load(frame *rtda.Frame, loader func(arr *heap.Object, idx int32)) {
	idx := frame.OperandStack().PopInt()
	arr := frame.OperandStack().PopRef()

	arr.AssertArrIdx(idx)

	loader(arr, idx)
}

func (*AALOAD) Execute(frame *rtda.Frame) {
	_load(
		frame,
		func(arr *heap.Object, idx int32) {
			frame.OperandStack().PushRef(arr.Refs()[idx])
		},
	)
}

func (*BALOAD) Execute(frame *rtda.Frame) {
	_load(
		frame,
		func(arr *heap.Object, idx int32) {
			frame.OperandStack().PushInt(int32(arr.Bytes()[idx]))
		},
	)
}

func (*CALOAD) Execute(frame *rtda.Frame) {
	_load(
		frame,
		func(arr *heap.Object, idx int32) {
			frame.OperandStack().PushInt(int32(arr.Chars()[idx]))
		},
	)
}

func (*DALOAD) Execute(frame *rtda.Frame) {
	_load(
		frame,
		func(arr *heap.Object, idx int32) {
			frame.OperandStack().PushDouble(arr.Doubles()[idx])
		},
	)
}

func (*FALOAD) Execute(frame *rtda.Frame) {
	_load(
		frame,
		func(arr *heap.Object, idx int32) {
			frame.OperandStack().PushFloat(arr.Floats()[idx])
		},
	)
}

func (*IALOAD) Execute(frame *rtda.Frame) {
	_load(
		frame,
		func(arr *heap.Object, idx int32) {
			frame.OperandStack().PushInt(arr.Ints()[idx])
		},
	)
}

func (*LALOAD) Execute(frame *rtda.Frame) {
	_load(
		frame,
		func(arr *heap.Object, idx int32) {
			frame.OperandStack().PushLong(arr.Longs()[idx])
		},
	)
}

func (*SALOAD) Execute(frame *rtda.Frame) {
	_load(
		frame,
		func(arr *heap.Object, idx int32) {
			frame.OperandStack().PushInt(int32(arr.Shorts()[idx]))
		},
	)
}
