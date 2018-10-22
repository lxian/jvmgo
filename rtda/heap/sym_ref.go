package heap

import "jvmgo/classfile"

type SymRef struct {
	constantPool *ConstantPool
	className string
	class *Class
}

type MemberRef struct {
	SymRef
	name string
	descriptor string
}

func (mem *MemberRef) copyFromMemberInfo(info *classfile.ConstantMemberRefInfo) {
	mem.className = info.ClassName()
	mem.name, mem.descriptor = info.NameAndType()
}

func (symRef *SymRef) ResolvedClass() {
}

func (symRef *SymRef) resolvedClassRef() {
	curClass := symRef.constantPool.class
	clz := curClass.classLoader.LoadClass(symRef.className)
	if
}
