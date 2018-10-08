package classfile

type ConstantClassInfo struct {
	constantPool ConstantPool
	nameIndex uint16
}

func (classInfo *ConstantClassInfo) readInfo(reader *ClassReader) {
	classInfo.nameIndex = reader.readUint16()
}

func (classInfo *ConstantClassInfo) Name() string {
	return classInfo.constantPool.getUtf8String(classInfo.nameIndex)
}

