package lang

import (
	"jvmgo/native"
	"jvmgo/rtda"
	"unsafe"
)

func init() {
	native.RegisterNativeMethod("java/lang/Object", "hashCode", "()I", hashCode)
}

func hashCode(frame *rtda.Frame) {
	this := frame.LocalVars().GetRef(0)
	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.OperandStack().PushInt(hash)
}
