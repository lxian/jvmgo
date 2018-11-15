package lang

import (
	"jvmgo/rtda"
	"jvmgo/native"
)

func init() {
	native.RegisterNativeMethod("java/lang/Throwable", "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

func fillInStackTrace(frame *rtda.Frame) {
	this := frame.LocalVars().GetRef(0)
	// TODO
	frame.OperandStack().PushRef(this)
}
