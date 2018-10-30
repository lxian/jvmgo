package heap

import "jvmgo/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (member *ClassMember) Name() string {
	return member.name
}

func (member *ClassMember) Descriptor() string {
	return member.descriptor
}

func (member *ClassMember) AccessFlags() uint16 {
	return member.accessFlags
}

func (member *ClassMember) Class() *Class {
	return member.class
}

func (member *ClassMember) copyInfoFromMemberInfo(memInfo *classfile.MemberInfo) {
	member.accessFlags = memInfo.AccessFlag()
	member.name = memInfo.Name()
	member.descriptor = memInfo.Descriptor()
}
