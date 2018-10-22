package heap

import "jvmgo/classfile"

type Field struct {
	ClassMember
	slotId uint
	constantValueIndex uint16
}

func newFields(class *Class, fieldsInfo []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(fieldsInfo))
	for i, fieldInfo := range fieldsInfo {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].constantValueIndex = fieldInfo.GetConstantValueIndex()
		fields[i].copyInfoFromMemberInfo(fieldInfo)
	}

	return fields
}
