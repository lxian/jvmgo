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

func (memInfo *MemberInfo) AccessFlag() uint16 {
	return memInfo.accessFlag
}

func (memInfo *MemberInfo) Name() string {
	return memInfo.constantPool.getUtf8String(memInfo.nameIndex)
}

func (memInfo *MemberInfo) Descriptor() string {
	return memInfo.constantPool.getUtf8String(memInfo.descriptorIndex)
}

func (memInfo *MemberInfo) FindCodeAttribute() *CodeAttribute {
	for _, attrInfo := range memInfo.attributes {
		switch attrInfo.(type) {
		case *CodeAttribute:
			return attrInfo.(*CodeAttribute)
		}
	}
	return nil
}

func (memInfo *MemberInfo) GetConstantValueAttribute() *ConstantValueAttribute {
	for _, attrInfo := range memInfo.attributes {
		switch attrInfo.(type) {
		case *ConstantValueAttribute:
			return attrInfo.(*ConstantValueAttribute)
		}
	}
	return nil
}
