package lang

import (
	"jvmgo/rtda"
	"jvmgo/native"
	"jvmgo/rtda/heap"
)

func init() {
	native.RegisterNativeMethod("java/lang/Class", "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.RegisterNativeMethod("java/lang/Class", "getName0", "()Ljava/lang/String;", getName0)
	native.RegisterNativeMethod("java/lang/Class", "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
}

func getPrimitiveClass(frame *rtda.Frame) {
	name := heap.GoString(frame.LocalVars().GetRef(0))
	loader := frame.Method().Class().ClassLoader()

	clz := loader.LoadClass(name)
	frame.OperandStack().PushRef(clz.JClassObj())
}

func getName0(frame *rtda.Frame) {
	this := frame.LocalVars().GetRef(0)

	clzName := heap.GetJavaClassName(this.Extra().(*heap.Class).Name())
	frame.OperandStack().PushRef(heap.JString(clzName, frame.Method().Class().ClassLoader()))
}

func desiredAssertionStatus0(frame *rtda.Frame) {
	frame.OperandStack().PushBoolean(false)
}