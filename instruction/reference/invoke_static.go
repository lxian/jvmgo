package reference

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

type INVOKE_STATIC struct {
	instruction.Index16Instruction
}

func (inst *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	invokerClass := frame.Method().Class()
	methodRef := frame.Method().Class().ConstantPool().GetConstant(inst.Index).(heap.MethodRef)

	method := methodRef.ResolvedMethod()
	class := methodRef.ResolvedClass()

	if !method.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
}
