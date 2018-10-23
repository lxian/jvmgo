package heap

import "jvmgo/classfile"

type Field struct {
	ClassMember
	slotId             uint
	constantValueIndex uint16
}

func (f *Field) SlotId() uint {
	return f.slotId
}

func newFields(class *Class, fieldsInfo []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(fieldsInfo))
	for i, fieldInfo := range fieldsInfo {
		fields[i] = &Field{}
		fields[i].class = class
		fields[i].copyAttributes(fieldInfo)
		fields[i].copyInfoFromMemberInfo(fieldInfo)
	}

	return fields
}

func (field *Field) copyAttributes(fieldInfo *classfile.MemberInfo) {
	if constAttrInfo := fieldInfo.GetConstantValueAttribute(); constAttrInfo != nil {
		field.constantValueIndex = constAttrInfo.ValueIndex()
	}
}

func (field *Field) isAccessibleTo(class *Class) bool {
	if HasFlag(field.accessFlags, ACC_PUBLIC) {
		return true
	}

	if HasFlag(field.accessFlags, ACC_PROTECTED) {
		return field.class == class || class.isSubClassOf(field.class) || class.packageName() == field.class.packageName()
	}

	if !HasFlag(field.accessFlags, ACC_PRIVATE) {
		return class.packageName() == field.class.packageName()
	}

	return class == field.class
}
