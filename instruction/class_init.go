package instruction

import (
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func InitClass(thread *rtda.Thread, class *heap.Class) {

	// mark class init started flag to true
	class.StartInit()

	// schedule init
	initMethod := class.FindInitMethod()
	if initMethod != nil {
		thread.PushFrame(rtda.NewFrame(thread, initMethod))
	}

	// init super
	if !class.IsInterface() {
		superClz := class.SuperClass()
		if superClz != nil && superClz.InitStarted() {
			InitClass(thread, superClz)
		}
	}
}
