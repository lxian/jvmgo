package classfile

type ConstantNameAndType struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (nameAndTypeInfo *ConstantNameAndType) readInfo(reader *ClassReader) {
	nameAndTypeInfo.nameIndex = reader.readUint16()
	nameAndTypeInfo.descriptorIndex = reader.readUint16()
}
