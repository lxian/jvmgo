package control

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
	"jvmgo/native/java/lang"
	"jvmgo/rtda/heap"
	"fmt"
	"os"
)

type ATHROW struct {
	instruction.NoOperandsInstruction
}

func (inst *ATHROW) Execute(frame *rtda.Frame) {

	expObj := frame.OperandStack().PopRef()
	if expObj == nil {
		panic("java/lang/NullPointerException")
	}
	for {
		method := frame.Method()
		thread := frame.Thread()
		pc := thread.PC()

		handlerPc := method.ExceptionTable().FindExceptionHandler(pc, expObj)
		if handlerPc != -1 {
			frame.SetNextPC(handlerPc)
			frame.OperandStack().Clear()
			frame.OperandStack().PushRef(expObj)
			return
		}

		if thread.IsStackEmpty() {
			jMsg := expObj.GetRefVar("detailMessage", "Ljava/lang/String;")
			goMsg := heap.GoString(jMsg)
			fmt.Fprintln(os.Stderr, expObj.Class().JavaName() + ": " + goMsg)

			traceEles := expObj.Extra().([]*lang.StackTraceElement)
			for _, traceEl := range traceEles {
				fmt.Fprintln(os.Stderr, "\tat " + traceEl.String())
			}
			return
		}
		thread.PopFrame()
	}
}
