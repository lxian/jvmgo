package classfile

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, constantPool ConstantPool) []AttributeInfo {
	attrCnt := reader.readUint16()
	attrs := make([]AttributeInfo, attrCnt)
	for i := range attrs {
		attrs[i] = readAttribute(reader, constantPool)
	}
	return attrs
}

func readAttribute(reader *ClassReader, constantPool ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrLen := reader.readUint32()
	attrName := constantPool.getUtf8String(attrNameIndex)
	attr := newAttributeInstance(attrName, attrLen, constantPool)
	attr.readInfo(reader)
	return attr
}

func newAttributeInstance(attrName string, attrLen uint32, constantPool ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{constantPool: constantPool}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Synthetic":
		return &SyntheticAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{constantPool: constantPool}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	default:
		return &UnparsedAttribute{name: attrName, length: attrLen}
	}
}
