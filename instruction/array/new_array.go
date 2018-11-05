package array

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

const (
	T_BOOLEAN = 4
	T_CHAR    = 5
	T_FLOAT   = 6
	T_DOUBLE  = 7
	T_BYTE    = 8
	T_SHORT   = 9
	T_INT     = 10
	T_LONG    = 11
)

type NEW_ARRAY struct {
	atype uint8
}

func (inst *NEW_ARRAY) FetchOperands(reader *instruction.ByteCodeReader) {
	inst.atype = reader.ReadUint8()
}

func (inst *NEW_ARRAY) Execute(frame *rtda.Frame) {
	eleCount := frame.OperandStack().PopInt()
	arrClz := arrayClassType(inst.atype, frame.Method().Class().ClassLoader())
	arr := arrClz.NewArray(uint(eleCount))
	frame.OperandStack().PushRef(arr)
}

func arrayClassType(atype uint8, loader *heap.ClassLoader) *heap.Class {
	switch atype {
	case T_BOOLEAN: return loader.LoadClass(heap.ARR_BOOL)
	case T_CHAR: return loader.LoadClass(heap.ARR_CHAR)
	case T_FLOAT: return loader.LoadClass(heap.ARR_FLOAT)
	case T_DOUBLE: return loader.LoadClass(heap.ARR_DOUBLE)
	case T_BYTE: return loader.LoadClass(heap.ARR_BYTE)
	case T_SHORT: return loader.LoadClass(heap.ARR_SHORT)
	case T_INT: return loader.LoadClass(heap.ARR_INT)
	case T_LONG: return loader.LoadClass(heap.ARR_LONG)
	}
	panic("Invalid atype "+string(atype))
}


