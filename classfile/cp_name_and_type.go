package classfile

type ConstantNameAndType struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (c *ConstantNameAndType) DescriptorIndex() uint16 {
	return c.descriptorIndex
}

func (c *ConstantNameAndType) NameIndex() uint16 {
	return c.nameIndex
}

func (nameAndTypeInfo *ConstantNameAndType) readInfo(reader *ClassReader) {
	nameAndTypeInfo.nameIndex = reader.readUint16()
	nameAndTypeInfo.descriptorIndex = reader.readUint16()
}
