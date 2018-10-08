package classfile

type ConstantValueAttribute struct {
	valueIndex uint16
}

func (attr *ConstantValueAttribute) readInfo(reader *ClassReader) {
	attr.valueIndex = reader.readUint16()
}
