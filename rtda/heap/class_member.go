package heap

import "jvmgo/classfile"

type ClassMember struct {
	accessFlags uint16
	name string
	descriptor string
	class *Class
}

func (member *ClassMember) copyInfoFromMemberInfo(memInfo *classfile.MemberInfo) {
	member.accessFlags = memInfo.AccessFlag()
	member.name = memInfo.Name()
	member.descriptor = memInfo.Descriptor()
}
