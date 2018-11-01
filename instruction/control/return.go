package control

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type RETURN struct {
	instruction.NoOperandsInstruction
}

func (*RETURN) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

type ARETURN struct {
	instruction.NoOperandsInstruction
}
type DRETURN struct {
	instruction.NoOperandsInstruction
}
type FRETURN struct {
	instruction.NoOperandsInstruction
}
type LRETURN struct {
	instruction.NoOperandsInstruction
}
type IRETURN struct {
	instruction.NoOperandsInstruction
}

func (*ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	poped := thread.PopFrame()
	invoker := thread.CurrentFrame()
	invoker.OperandStack().PushRef(poped.OperandStack().PopRef())
}

func (*DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	poped := thread.PopFrame()
	invoker := thread.CurrentFrame()
	invoker.OperandStack().PushDouble(poped.OperandStack().PopDouble())
}

func (*FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	poped := thread.PopFrame()
	invoker := thread.CurrentFrame()
	invoker.OperandStack().PushFloat(poped.OperandStack().PopFloat())
}

func (*LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	poped := thread.PopFrame()
	invoker := thread.CurrentFrame()
	invoker.OperandStack().PushLong(poped.OperandStack().PopLong())
}

func (*IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	poped := thread.PopFrame()
	invoker := thread.CurrentFrame()
	invoker.OperandStack().PushInt(poped.OperandStack().PopInt())
}
