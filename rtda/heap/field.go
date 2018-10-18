package heap

import "jvmgo/classfile"

type Field struct {
	ClassMember
}

func newFileds(class *Class, fieldsInfo []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(fieldsInfo))
	for i, fieldInfo := range fieldsInfo {
		fields[i] = &Field{ClassMember{class:class}}
		fields[i].copyInfoFromMemberInfo(fieldInfo)
	}

	return fields
}
