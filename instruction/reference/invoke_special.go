package reference

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

type INVOKE_SPECIAL struct {
	instruction.Index16Instruction
}

func (inst *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(inst.Index).(*heap.MethodRef)
	println(methodRef)
	frame.OperandStack().PopRef()
}

