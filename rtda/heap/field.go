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
		fields[i].copyAttributes(fieldInfo);
		fields[i].copyInfoFromMemberInfo(fieldInfo)
	}

	return fields
}

func (field *Field) copyAttributes(fieldInfo *classfile.MemberInfo) {
	if constAttrInfo := fieldInfo.GetConstantValueAttribute(); constAttrInfo != nil {
		field.constantValueIndex = constAttrInfo.ValueIndex()
	}
}
