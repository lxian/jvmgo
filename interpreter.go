package main

import (
	"fmt"
	"jvmgo/classfile"
	"jvmgo/instruction"
	"jvmgo/instruction/factory"
	"jvmgo/rtda"
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
			fmt.Printf("LocalVars: %v\n", frame.LocalVars())
			fmt.Printf("OperandStack: %v\n", frame.OperandStack())
			panic(r)
		}
	}(frame)
	loop(thread, bytecode)
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &instruction.ByteCodeReader{}

	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		instruction := factory.NewInstruction(opcode)
		instruction.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		fmt.Printf("pc: %2d inst: %T %v \n", pc, instruction, instruction)
		fmt.Printf("Locals: %v\n", frame.LocalVars())
		fmt.Printf("Stack: %v\n", frame.OperandStack())
		instruction.Execute(frame)
		fmt.Printf(">Locals: %v\n", frame.LocalVars())
		fmt.Printf(">Stack: %v\n", frame.OperandStack())
		fmt.Println("-----------")
	}
}
