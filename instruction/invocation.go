package instruction

import (
	"jvmgo/rtda/heap"
	"jvmgo/rtda"
)

func Invoke(method *heap.Method, frame *rtda.Frame) {
	thread := frame.Thread()
	newFrame := rtda.NewFrame(frame.Thread(), method)

	for i := int(method.ArgsSlotCount()-1); i>=0; i-- {
		slot := frame.OperandStack().PopSlot()
		newFrame.LocalVars().SetSlot(uint(i), slot)
	}

	thread.PushFrame(newFrame)
}

func AssertThisRef(method *heap.Method, frame *rtda.Frame) {
	if FindThisRef(method, frame) == nil {
		panic("java.lang.IncompatibleClassChangeError")
	}
}

func FindThisRef(method *heap.Method, frame *rtda.Frame) *heap.Object {
	if method.IsStatic() {
		panic("No implicit this ref for static method")
	}

	return frame.OperandStack().PeekSlotBelow(method.ArgsSlotCount()-1).Ref()
}
