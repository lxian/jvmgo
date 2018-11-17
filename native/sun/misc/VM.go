package misc

import (
	"jvmgo/instruction"
	"jvmgo/native"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func init() {
	native.RegisterNativeMethod("sun/misc/VM", "initialize", "()V", initialize)
}

func initialize(frame *rtda.Frame) {
	vmClz := frame.Method().Class()
	savedProps := vmClz.FindStaticVarRef("savedProps", "Ljava/util/Properties;")

	setPropMethod := savedProps.GetObjectMethod(
		"setProperty",
		"(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	frame.OperandStack().PushRef(savedProps)
	frame.OperandStack().PushRef(heap.JString("foo", frame.Method().Class().ClassLoader()))
	frame.OperandStack().PushRef(heap.JString("bar", frame.Method().Class().ClassLoader()))
	instruction.Invoke(setPropMethod, frame)

	//loader := frame.Method().Class().ClassLoader()
	//sysClz := loader.LoadClass("java/lang/System")
	//method := sysClz.FindStaticMethod("initializeSystemClass", "()V")
	//instruction.Invoke(method, frame)
}
