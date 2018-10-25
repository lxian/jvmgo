package heap

import "jvmgo/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

func newFieldRef(constantPool *ConstantPool, info *classfile.ConstantFieldRefInfo) *FieldRef {
	fieldRef := &FieldRef{}
	fieldRef.constantPool = constantPool
	fieldRef.copyFromMemberInfo(&info.ConstantMemberRefInfo)
	return fieldRef
}

func (fieldRef *FieldRef) ResolvedField() *Field {
	if fieldRef.field == nil {
		fieldRef.resolveField()
	}
	return fieldRef.field
}

func (fieldRef *FieldRef) resolveField() {
	curClz := fieldRef.constantPool.class
	fieldClz := fieldRef.ResolvedClass()
	field := lookupFiled(fieldClz, fieldRef.name, fieldRef.descriptor)

	if !field.isAccessibleTo(curClz) {
		panic("Illegal access error")
	}
	fieldRef.field = field
}

func lookupFiled(class *Class, name string, desc string) *Field {
	for _, filed := range class.fields {
		if filed.name == name && filed.descriptor == desc {
			return filed
		}
	}

	for _, iface := range class.interfaces {
		if f := lookupFiled(iface, name, desc); f != nil {
			return f
		}
	}

	if class.superClass != nil {
		if f := lookupFiled(class.superClass, name, desc); f != nil {
			return f
		}
	}

	panic("No such field error")
}
