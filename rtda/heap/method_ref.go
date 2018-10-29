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

func (methodRef *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if methodRef.method == nil {
		resolveInterfaceMethod(methodRef)
	}
	return methodRef.method
}

func resolveInterfaceMethod(methodRef *InterfaceMethodRef) {
	clz := methodRef.ResolvedClass()

	if !clz.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupInterfaceMethod(clz, methodRef.name, methodRef.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	if !method.IsAccessibleTo(methodRef.class) {
		panic("java.lang.IllegalAccessError")
	}

	methodRef.method = method
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

func (methodRef *MethodRef) ResolvedMethod() *Method {
	if methodRef.method == nil {
		resolveMethod(methodRef)
	}
	return methodRef.method
}

func resolveMethod(methodRef *MethodRef) {
	clz := methodRef.ResolvedClass()

	if clz.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupMethod(clz, methodRef.name, methodRef.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	if !method.IsAccessibleTo(methodRef.class) {
		panic("java.lang.IllegalAccessError")
	}

	methodRef.method = method
}

