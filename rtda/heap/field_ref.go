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

