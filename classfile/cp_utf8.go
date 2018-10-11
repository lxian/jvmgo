package classfile

// UTF8
type ConstantUtf8Info struct {
	value string
}

func (utf8Info *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	utf8Info.value = string(reader.readBytes(length))
}

// String
type ConstantStringInfo struct {
	constantPool ConstantPool
	utf8Index    uint16
}

func (stringInfo *ConstantStringInfo) readInfo(reader *ClassReader) {
	stringInfo.utf8Index = reader.readUint16()
}

func (stringInfo *ConstantStringInfo) String() string {
	return stringInfo.constantPool.getUtf8String(stringInfo.utf8Index)
}
