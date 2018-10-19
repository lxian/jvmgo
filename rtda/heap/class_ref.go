package heap

import "jvmgo/classfile"

type ClassRef struct {
	SymRef
}

func newClassRef(constantPool *ConstantPool, info *classfile.ConstantClassInfo) *ClassRef {
	return &ClassRef{SymRef{constantPool:constantPool, className:info.Name()}}
}

