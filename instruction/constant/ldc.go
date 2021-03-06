package constant

import (
	"fmt"
	"jvmgo/instruction"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func ldc(idx uint, frame *rtda.Frame) {
	val := frame.Method().Class().ConstantPool().GetConstant(idx)
	switch val.(type) {
	case int32:
		frame.OperandStack().PushInt(val.(int32))
	case float32:
		frame.OperandStack().PushFloat(val.(float32))
	case string:
		frame.OperandStack().PushRef(heap.InternedJString(val.(string), frame.Method().Class().ClassLoader()))
	case *heap.ClassRef:
		frame.OperandStack().PushRef(val.(*heap.ClassRef).ResolvedClass().JClassObj())
	default:
		panic(fmt.Sprintf("todo ldc %v", val))
	}
}

type LDC struct {
	instruction.Index8Instruction
}

func (inst *LDC) Execute(frame *rtda.Frame) {
	ldc(inst.Index, frame)
}

type LDC_W struct {
	instruction.Index16Instruction
}

func (inst *LDC_W) Execute(frame *rtda.Frame) {
	ldc(inst.Index, frame)
}

type LDC2_W struct {
	instruction.Index16Instruction
}

func (inst *LDC2_W) Execute(frame *rtda.Frame) {
	val := frame.Method().Class().ConstantPool().GetConstant(inst.Index)
	switch val.(type) {
	case float64:
		frame.OperandStack().PushDouble(val.(float64))
	case int64:
		frame.OperandStack().PushLong(val.(int64))
	default:
		panic(fmt.Sprintf("todo ldc %v", val))
	}
}
