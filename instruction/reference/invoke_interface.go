package reference

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

type INVOKE_INTERFACE struct {
	idx uint
}

func (inst *INVOKE_INTERFACE) FetchOperands(reader *instruction.ByteCodeReader) {
	inst.idx = uint(reader.ReadUint16())
	reader.ReadUint8() // skip the count byte
	reader.ReadUint8() // skip the 0 byte
}

func (inst *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(inst.idx).(*heap.MethodRef)
	method := methodRef.ResolvedMethod()

	if method.IsStatic() || method.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	obj := instruction.FindThisRef(method, frame)
	instruction.AssertThisRef(method, frame)

	if !obj.Class().Implemented(method.Class()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	toInvoke := heap.LookupMethodInClass(obj.Class(), method.Name(), method.Descriptor())
	if toInvoke == nil || toInvoke.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	if !toInvoke.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}
	instruction.Invoke(toInvoke, frame)
}
