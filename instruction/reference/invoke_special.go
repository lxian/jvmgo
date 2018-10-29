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

	method := methodRef.ResolvedMethod()
	invokerClz := frame.Method().Class()
	resolvedClz := methodRef.ResolvedClass()

	if method.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if method.Name() == "<init>" || method.Class() != resolvedClz {
		panic("java.lang.NoSuchMethodError")
	}


	if heap.HasFlag(invokerClz.AccessFlags(), heap.ACC_SUPER) &&
		invokerClz.IsSubClassOf(resolvedClz) &&
			method.Name() != "<init>" {
				method = heap.LookupMethodInClass(invokerClz, method.Name(), method.Descriptor())
	}

	if method == nil || heap.HasFlag(method.AccessFlags(), heap.ACC_ABSTRACT) {
		panic("java.lang.AbstractMethodError")
	}

	instruction.Invoke(method, frame)
}

