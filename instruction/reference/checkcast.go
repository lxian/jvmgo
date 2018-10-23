package reference

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

type CHECK_CAST struct {
	instruction.Index16Instruction
}

func (inst *CHECK_CAST) Execute(frame *rtda.Frame) {
	clzCastToRef := frame.Method().Class().ConstantPool().GetConstant(inst.Index).(*heap.ClassRef)
	clzCastTo := clzCastToRef.ResolvedClass()
	obj := frame.OperandStack().PeekSlot().Ref()
	if obj == nil {
		return
	}

	if !obj.IsInstanceOf(clzCastTo) {
		panic("Class Cast Exception")
	}
}
