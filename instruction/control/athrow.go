package control

import (
	"fmt"
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type ATHROW struct {
	instruction.NoOperandsInstruction
}

func (*ATHROW) Execute(frame *rtda.Frame) {

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
			fmt.Println("Cannnot handle the exception")
			return // TODO print somehting
		}
		thread.PopFrame()
	}
}
