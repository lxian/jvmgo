package classfile

type ConstantValueAttribute struct {
	valueIndex uint16
}

func (attr *ConstantValueAttribute) readInfo(reader *ClassReader) {
	attr.valueIndex = reader.readUint16()
}

func (attr *ConstantValueAttribute) ValueIndex() uint16 {
	return attr.valueIndex
}
