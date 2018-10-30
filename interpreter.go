package main

import (
	"fmt"
	"jvmgo/instruction"
	"jvmgo/instruction/factory"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
	"strings"
)

func interpret(method *heap.Method, logInst bool) {
	thread := rtda.NewThread()
	frame := rtda.NewFrame(thread, method)
	thread.PushFrame(frame)

	defer func(frm *rtda.Frame) {
		if r := recover(); r != nil {
			logFrames(frm.Thread())
			panic(r)
		}
	}(frame)
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
	fmt.Printf("%v.%v() pc %2d inst %T %v \n", className, methodName, frame.Thread().PC(), inst, inst)
}

func logFrames(thread *rtda.Thread) {
	fmt.Printf("LocalVars: %v\n", thread.CurrentFrame().LocalVars())
	fmt.Printf("OperandStack: %v\n\n", thread.CurrentFrame().OperandStack())
	i := 0
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		class := method.Class()

		fmt.Printf(strings.Repeat("  ", i))
		fmt.Printf("pc:%4d %v.%v%v\n", frame.NextPC(), class.Name(), method.Name(), method.Descriptor())
		i += 1
	}
}
