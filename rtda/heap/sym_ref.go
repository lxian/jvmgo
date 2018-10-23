package heap

import "jvmgo/classfile"

type SymRef struct {
	constantPool *ConstantPool
	className    string
	class        *Class
}

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

func (mem *MemberRef) copyFromMemberInfo(info *classfile.ConstantMemberRefInfo) {
	mem.className = info.ClassName()
	mem.name, mem.descriptor = info.NameAndType()
}

func (symRef *SymRef) ResolvedClass() *Class {
	if symRef.class == nil {
		symRef.resolveClassRef()
	}
	return symRef.class
}

func (symRef *SymRef) resolveClassRef() {
	curClass := symRef.constantPool.class
	clz := curClass.classLoader.LoadClass(symRef.className)
	if !clz.isAccessibleTo(curClass) {
		panic("Illegal access error")
	}
	symRef.class = clz
}
