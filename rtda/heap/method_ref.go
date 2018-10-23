package heap

import "jvmgo/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method *Method
}

func newInterfaceMethodRef(constantPool *ConstantPool, info *classfile.ConstantInterfaceMethodRefInfo) *InterfaceMethodRef {
	interfaceMethodRef := &InterfaceMethodRef{}
	interfaceMethodRef.constantPool = constantPool
	interfaceMethodRef.copyFromMemberInfo(&info.ConstantMemberRefInfo)
	return interfaceMethodRef
}

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(constantPool *ConstantPool, info *classfile.ConstantMethodRefInfo) *MethodRef {
	methodRef := &MethodRef{}
	methodRef.constantPool = constantPool
	methodRef.copyFromMemberInfo(&info.ConstantMemberRefInfo)
	return methodRef
}
