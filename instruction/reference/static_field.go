package reference

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func getStaticFieldAndClz(idx uint, frame *rtda.Frame) (*heap.Field, *heap.Class) {
	cp := frame.Method().Class().ConstantPool()
	fieldRef := cp.GetConstant(idx).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	clz := fieldRef.ResolvedClass()

	if !heap.HasFlag(field.AccessFlags(), heap.ACC_STATIC) {
		panic("Static operation on non static field")
	}

	return field, clz
}

type PUT_STATIC struct{ instruction.Index16Instruction }

func (inst *PUT_STATIC) Execute(frame *rtda.Frame) {
	field, clz := getStaticFieldAndClz(inst.Index, frame)

	if !clz.InitStarted() {
		frame.RevertNextPC()
		instruction.InitClass(frame.Thread(), clz)
		return
	}

	slotId := field.SlotId()
	switch field.Descriptor()[0] {
	case heap.BOOL, heap.SHORT, heap.CHAR, heap.INT:
		clz.StaticVars().SetInt(slotId, frame.OperandStack().PopInt())
	case heap.LONG:
		clz.StaticVars().SetLong(slotId, frame.OperandStack().PopLong())
	case heap.FLOAT:
		clz.StaticVars().SetFloat(slotId, frame.OperandStack().PopFloat())
	case heap.DOUBLE:
		clz.StaticVars().SetDouble(slotId, frame.OperandStack().PopDouble())
	case heap.OBJECT, heap.ARRAY:
		clz.StaticVars().SetRef(slotId, frame.OperandStack().PopRef())
	}
}

type GET_STATIC struct{ instruction.Index16Instruction }

func (inst *GET_STATIC) Execute(frame *rtda.Frame) {
	field, clz := getStaticFieldAndClz(inst.Index, frame)

	if !clz.InitStarted() {
		frame.RevertNextPC()
		instruction.InitClass(frame.Thread(), clz)
		return
	}

	slotId := field.SlotId()
	switch field.Descriptor()[0] {
	case heap.BOOL, heap.SHORT, heap.CHAR, heap.INT:
		frame.OperandStack().PushInt(clz.StaticVars().GetInt(slotId))
	case heap.LONG:
		frame.OperandStack().PushLong(clz.StaticVars().GetLong(slotId))
	case heap.FLOAT:
		frame.OperandStack().PushFloat(clz.StaticVars().GetFloat(slotId))
	case heap.DOUBLE:
		frame.OperandStack().PushDouble(clz.StaticVars().GetDouble(slotId))
	case heap.OBJECT, heap.ARRAY:
		frame.OperandStack().PushRef(clz.StaticVars().GetRef(slotId))
	}
}
