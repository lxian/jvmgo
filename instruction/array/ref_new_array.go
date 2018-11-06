package array

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

type ANEW_ARRAY struct {
	instruction.Index16Instruction
}

func (inst *ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	eleClzRef := cp.GetConstant(inst.Index).(*heap.ClassRef)
	eleClz := eleClzRef.ResolvedClass()
	arrClz := eleClz.ArrayClass()

	eleCount := frame.OperandStack().PopInt()
	arr := arrClz.NewArray(uint(eleCount))

	frame.OperandStack().PushRef(arr)
}
