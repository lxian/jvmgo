package reference

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

type INSTANCE_OF struct {
	instruction.Index16Instruction
}

func (inst *INSTANCE_OF) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(doCheckInstanceOf(inst, frame))
}

func doCheckInstanceOf(inst *INSTANCE_OF, frame *rtda.Frame) int32 {
	const (
		TRUE  = 1
		FALSE = 0
	)
	obj := frame.OperandStack().PopRef()
	if obj == nil {
		return TRUE
	}

	classRef := frame.Method().Class().ConstantPool().GetConstant(inst.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()

	if obj.IsInstanceOf(class) {
		return TRUE
	} else {
		return FALSE
	}

}
