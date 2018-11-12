package lang

import (
	"jvmgo/native"
	"jvmgo/rtda"
	"unsafe"
)

func init() {
	native.RegisterNativeMethod("java/lang/Object", "hashCode", "()I", hashCode)
	native.RegisterNativeMethod("java/lang/Object", "getClass", "()Ljava/lang/Class;", hashCode)
}

func hashCode(frame *rtda.Frame) {
	this := frame.LocalVars().GetRef(0)
	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}

func getClass(frame *rtda.Frame) {
	this := frame.LocalVars().GetRef(0)
	jClzObj := this.Class().JClassObj()
	frame.OperandStack().PushRef(jClzObj)
}
