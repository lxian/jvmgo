package classfile

type MemberInfo struct {
	constantPool    ConstantPool
	accessFlag      uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, constantPool ConstantPool) []*MemberInfo {
	memberCount := uint16(reader.readUint16())
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, constantPool)
	}
	return members
}

func readMember(reader *ClassReader, constantPool ConstantPool) *MemberInfo {
	return &MemberInfo{
		constantPool:    constantPool,
		accessFlag:      reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, constantPool),
	}
}

func (memInfo *MemberInfo) FindCodeAttribute() *CodeAttribute {
	for _, attrInfo := range memInfo.attributes {
		switch attrInfo.(type) {
		case CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

