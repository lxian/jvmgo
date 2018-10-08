package classfile

type SourceFileAttribute struct {
	constantPool ConstantPool
	sourceFileIndex uint16
}

func (attr *SourceFileAttribute) readInfo(reader *ClassReader) {
	attr.sourceFileIndex = reader.readUint16()
}

func (attr *SourceFileAttribute) SourceFileName() string {
	return attr.constantPool.getUtf8String(attr.sourceFileIndex)
}