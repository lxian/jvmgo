package main

import (
	"fmt"
	"jvmgo/instruction"
	"jvmgo/instruction/factory"
	"jvmgo/rtda"
	"strings"
)

func interpret(thread *rtda.Thread, logInst bool) {
	defer func(thd *rtda.Thread) {
		if r := recover(); r != nil {
			logFrames(thd)
			panic(r)
		}
	}(thread)
	loop(thread, logInst)
}

func loop(thread *rtda.Thread, logInst bool) {
	reader := &instruction.ByteCodeReader{}

	for {
		frame := thread.CurrentFrame()
		if frame == nil {
			break
		}

		pc := frame.NextPC()
		thread.SetPC(pc) // thread keep the PC of current op

		reader.Reset(frame.Method().Code(), pc) // reader keep the PC of current reading byte
		opcode := reader.ReadUint8()
		inst := factory.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC()) // frame keep the PC of current execution/jump

		if logInst {
			logInstruction(inst, frame)
		}

		inst.Execute(frame)
	}
}

func logInstruction(inst instruction.Instruction, frame *rtda.Frame) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	fmt.Printf(strings.Repeat(".", int(frame.Thread().StackSize())))
	fmt.Printf("%v.%v() pc %2d inst %T %v \n", className, methodName, frame.Thread().PC(), inst, inst)
}

func logFrames(thread *rtda.Thread) {
	fmt.Printf("LocalVars: %v\n", thread.CurrentFrame().LocalVars())
	fmt.Printf("OperandStack: %v\n\n", thread.CurrentFrame().OperandStack())
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		class := method.Class()

		fmt.Printf(">pc:%4d %v.%v%v\n", frame.NextPC(), class.Name(), method.Name(), method.Descriptor())
		fmt.Printf("%v \n", frame)
	}
}
