package heap

import "jvmgo/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (c *ClassMember) Descriptor() string {
	return c.descriptor
}

func (c *ClassMember) AccessFlags() uint16 {
	return c.accessFlags
}

func (c *ClassMember) Class() *Class {
	return c.class
}

func (member *ClassMember) copyInfoFromMemberInfo(memInfo *classfile.MemberInfo) {
	member.accessFlags = memInfo.AccessFlag()
	member.name = memInfo.Name()
	member.descriptor = memInfo.Descriptor()
}
