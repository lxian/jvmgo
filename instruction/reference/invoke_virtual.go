package reference

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
	"fmt"
)

type INVOKE_VIRTUAL struct {
	instruction.Index16Instruction
}

func (inst *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	methodRef := cp.GetConstant(inst.Index).(*heap.MethodRef)
	method := methodRef.ResolvedMethod()

	if method.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	obj := instruction.FindThisRef(method, frame)
	if checkAndHackPrintln(obj, method, frame) {
		return
	}
	instruction.AssertThisRef(method, frame)

	toInvoke := heap.LookupMethodInClass(obj.Class(), method.Name(), method.Descriptor())
	if toInvoke == nil || toInvoke.IsAbstract()  {
		panic("java.lang.AbstractMethodError")
	}
	instruction.Invoke(toInvoke, frame)
}

func checkAndHackPrintln(obj *heap.Object, method *heap.Method, frame *rtda.Frame) bool {
	if obj == nil {
		if method.Name() == "println" {
			switch method.Descriptor() {
			case "(Z)V": fmt.Printf("%v\n", frame.OperandStack().PopInt() != 0)
			case "(C)V": fmt.Printf("%c\n", frame.OperandStack().PopInt())
			case "(S)V": fmt.Printf("%v\n", frame.OperandStack().PopInt())
			case "(B)V": fmt.Printf("%v\n", frame.OperandStack().PopInt())
			case "(I)V": fmt.Printf("%v\n", frame.OperandStack().PopInt())
			case "(J)V": fmt.Printf("%v\n", frame.OperandStack().PopLong())
			case "(D)V": fmt.Printf("%v\n", frame.OperandStack().PopDouble())
			case "(F)V": fmt.Printf("%v\n", frame.OperandStack().PopFloat())
			default: panic("println: " +method.Descriptor())
			}
			frame.OperandStack().PopRef() // pop out the nil ref
			return true
		}
	}
	return false
}
