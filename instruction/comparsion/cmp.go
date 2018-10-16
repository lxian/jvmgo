package comparsion

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type LCMP struct {
	instruction.NoOperandsInstruction
}

func (*LCMP) Execute(frame *rtda.Frame) {
	v2 := frame.OperandStack().PopLong()
	v1 := frame.OperandStack().PopLong()
	if v1 > v2 {
		frame.OperandStack().PushInt(1)
	} else if v1 < v2 {
		frame.OperandStack().PushInt(-1)
	} else {
		frame.OperandStack().PushInt(0)
	}
}

type FCMPG struct {
	instruction.NoOperandsInstruction
}
type FCMPL struct {
	instruction.NoOperandsInstruction
}

func (*FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

func (*FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}

func _fcmp(frame *rtda.Frame, gFlag bool) {
	v2 := frame.OperandStack().PopFloat()
	v1 := frame.OperandStack().PopFloat()
	if v1 > v2 {
		frame.OperandStack().PushInt(1)
	} else if v1 < v2 {
		frame.OperandStack().PushInt(-1)
	} else if v1 == v2 {
		frame.OperandStack().PushInt(0)
	} else if gFlag {
		frame.OperandStack().PushInt(1)
	} else {
		frame.OperandStack().PushInt(-1)
	}
}

type DCMPG struct {
	instruction.NoOperandsInstruction
}
type DCMPL struct {
	instruction.NoOperandsInstruction
}

func (*DCMPG) Execute(frame *rtda.Frame) {
	_dcmp(frame, true)
}

func (*DCMPL) Execute(frame *rtda.Frame) {
	_dcmp(frame, false)
}

func _dcmp(frame *rtda.Frame, gFlag bool) {
	v2 := frame.OperandStack().PopDouble()
	v1 := frame.OperandStack().PopDouble()
	if v1 > v2 {
		frame.OperandStack().PushInt(1)
	} else if v1 < v2 {
		frame.OperandStack().PushInt(-1)
	} else if v1 == v2 {
		frame.OperandStack().PushInt(0)
	} else if gFlag {
		frame.OperandStack().PushInt(1)
	} else {
		frame.OperandStack().PushInt(-1)
	}
}
