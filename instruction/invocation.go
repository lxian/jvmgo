package instruction

import (
	"jvmgo/rtda/heap"
	"jvmgo/rtda"
)

func Invoke(method *heap.Method, frame *rtda.Frame) {
	thread := frame.Thread()
	newFrame := rtda.NewFrame(frame.Thread(), method)
	thread.PushFrame(newFrame)

	for i := method.ArgsCount()-1; i>=0; i-- {
		slot := frame.OperandStack().PopSlot()
		newFrame.LocalVars().SetSlot(i, slot)
	}
}
