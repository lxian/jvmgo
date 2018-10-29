package instruction

import (
	"jvmgo/rtda/heap"
	"jvmgo/rtda"
)

func Invoke(method *heap.Method, frame *rtda.Frame) {
	invoke(method, frame, false)
}

func InvokeStatic(method *heap.Method, frame *rtda.Frame) {
	invoke(method, frame, true)
}

func invoke(method *heap.Method, frame *rtda.Frame, static bool) {
	thread := frame.Thread()
	newFrame := rtda.NewFrame(frame.Thread(), method)

	for i := method.ArgsCount()-1; i>=0; i-- {
		slot := frame.OperandStack().PopSlot()
		newFrame.LocalVars().SetSlot(i, slot)
	}

	if !static && newFrame.LocalVars().GetRef(0) == nil {
		panic("java.lang.NullPointerException")
	}

	thread.PushFrame(newFrame)
}
