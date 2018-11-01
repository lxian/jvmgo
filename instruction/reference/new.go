package reference

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

type NEW struct {
	instruction.Index16Instruction
}

func (inst *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(inst.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	if !class.InitStarted() {
		frame.RevertNextPC()
		instruction.InitClass(frame.Thread(), class)
		return
	}

	if heap.HasFlag(class.AccessFlags(), heap.ACC_INTERFACE) || heap.HasFlag(class.AccessFlags(), heap.ACC_ABSTRACT) {
		panic("Instantiation error")
	}

	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
}
