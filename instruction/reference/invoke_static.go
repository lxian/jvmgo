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
	methodRef := frame.Method().Class().ConstantPool().GetConstant(inst.Index).(*heap.MethodRef)
	method := methodRef.ResolvedMethod()
	class := methodRef.ResolvedClass()

	if !class.InitStarted() {
		frame.RevertNextPC()
		instruction.InitClass(frame.Thread(), class)
		return
	}


	// skip native methods register for now
	if method.Name() == "registerNatives" && method.Descriptor() == "()V" {
		frame.Thread().PopFrame()
		return
	}
	instruction.Invoke(method, frame)
}
