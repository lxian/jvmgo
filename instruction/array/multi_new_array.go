package array

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

type MULTI_ANEW_ARRAY struct {
	typeIdx      uint
	dimensionCnt uint8
}

func (inst *MULTI_ANEW_ARRAY) FetchOperands(reader *instruction.ByteCodeReader) {
	inst.typeIdx = uint(reader.ReadUint16())
	inst.dimensionCnt = reader.ReadUint8()
}

func (inst *MULTI_ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	clzRef := cp.GetConstant(inst.typeIdx).(*heap.ClassRef)
	resolvedClz := clzRef.ResolvedClass()

	dims := make([]int32, inst.dimensionCnt)
	for i := int(inst.dimensionCnt - 1); i >= 0; i-- {
		dims[i] = frame.OperandStack().PopInt()
	}

	arr := newMultiArray(resolvedClz, dims)
	frame.OperandStack().PushRef(arr)
}

func newMultiArray(arrClz *heap.Class, dims []int32) *heap.Object {
	if len(dims) == 1 {
		return arrClz.NewArray(uint(dims[0]))
	}

	arr := arrClz.NewArray(uint(dims[0]))
	for i := range arr.Refs() {
		arr.Refs()[i] = newMultiArray(arrClz.ComponentClass(), dims[1:])
	}
	return arr
}
