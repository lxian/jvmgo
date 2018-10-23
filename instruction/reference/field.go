package reference

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func validateObjField(object *heap.Object, class *heap.Class, field *heap.Field) {
	if object == nil {
		panic("Null pointer exception")
	}
	if class != object.Class() {
		panic("Get Field reference to a different class")
	}
	if heap.HasFlag(field.AccessFlags(), heap.ACC_STATIC) {
		panic("Non static operation on static field")
	}
}

type PUT_FIELD struct {
	instruction.Index16Instruction
}

func (inst *PUT_FIELD) Execute(frame *rtda.Frame) {
	fieldRef := frame.Method().Class().ConstantPool().GetConstant(inst.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	clz := fieldRef.ResolvedClass()

	slotId := field.SlotId()
	switch field.Descriptor()[0] {
	case heap.BOOL, heap.SHORT, heap.CHAR, heap.INT:
		val := frame.OperandStack().PopInt()
		object := frame.OperandStack().PopRef()
		validateObjField(object, clz, field)
		object.Fields().SetInt(slotId, val)
	case heap.LONG:
		val := frame.OperandStack().PopLong()
		object := frame.OperandStack().PopRef()
		validateObjField(object, clz, field)
		object.Fields().SetLong(slotId, val)
	case heap.FLOAT:
		val := frame.OperandStack().PopFloat()
		object := frame.OperandStack().PopRef()
		validateObjField(object, clz, field)
		object.Fields().SetFloat(slotId, val)
	case heap.DOUBLE:
		val := frame.OperandStack().PopDouble()
		object := frame.OperandStack().PopRef()
		validateObjField(object, clz, field)
		object.Fields().SetDouble(slotId, val)
	case heap.OBJECT, heap.ARRAY:
		val := frame.OperandStack().PopRef()
		object := frame.OperandStack().PopRef()
		validateObjField(object, clz, field)
		object.Fields().SetRef(slotId, val)
	}
}

type GET_FIELD struct {
	instruction.Index16Instruction
}

func (inst *GET_FIELD) Execute(frame *rtda.Frame) {
	object := frame.OperandStack().PopRef()
	fieldRef := frame.Method().Class().ConstantPool().GetConstant(inst.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	clz := fieldRef.ResolvedClass()

	validateObjField(object, clz, field)

	slotId := field.SlotId()
	switch field.Descriptor()[0] {
	case heap.BOOL, heap.SHORT, heap.CHAR, heap.INT:
		frame.OperandStack().PushInt(object.Fields().GetInt(slotId))
	case heap.LONG:
		frame.OperandStack().PushLong(object.Fields().GetLong(slotId))
	case heap.FLOAT:
		frame.OperandStack().PushFloat(object.Fields().GetFloat(slotId))
	case heap.DOUBLE:
		frame.OperandStack().PushDouble(object.Fields().GetDouble(slotId))
	case heap.OBJECT, heap.ARRAY:
		frame.OperandStack().PushRef(object.Fields().GetRef(slotId))
	}
}
