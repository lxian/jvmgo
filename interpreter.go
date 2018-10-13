package main

import (
	"jvmgo/classfile"
	"jvmgo/rtda"
	"jvmgo/instruction"
	"fmt"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.FindCodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	bytecode := codeAttr.Code()

	thread := rtda.NewThread()
	frame := rtda.NewFrame(uint(maxLocals), uint(maxStack), thread)
	thread.PushFrame(frame)

	defer func(frame2 *rtda.Frame) {
		if r := recover(); r != nil {
			fmt.Printf("LocalVars: $v\n", frame.LocalVars())
			fmt.Printf("OperandStack: $v\n", frame.OperandStack())
			panic(r)
		}
	}(frame)
	loop(thread, bytecode)
}

func loop(thread *rtda.Thread, bytes []byte) {

}
