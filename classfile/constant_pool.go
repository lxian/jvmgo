package classfile

import "fmt"

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	length := int(reader.readUint16())
	var constantPool ConstantPool = make([]ConstantInfo, length)
	for i := 1; i < length; {
		constantPool[i] = readConstantInfo(reader, constantPool)

		switch constantPool[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i += 2
		default:
			i += 1
		}
	}
	return constantPool
}

func (constantPool ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := constantPool[index]; cpInfo != nil {
		return cpInfo
	}
	panic(fmt.Sprintf("Invalid Constant Pool Index %d", index))
}

func (constantPool ConstantPool) getUtf8String(index uint16) string {
	return constantPool.getConstantInfo(index).(*ConstantUtf8Info).value
}

func (constantPool ConstantPool) getClassName(index uint16) string {
	return constantPool.getConstantInfo(index).(*ConstantClassInfo).Name()
}

func (constantPool ConstantPool) getNameAndType(index uint16) (string, string) {
	nameAndType := constantPool.getConstantInfo(index).(*ConstantNameAndType)
	return constantPool.getUtf8String(nameAndType.nameIndex), constantPool.getUtf8String(nameAndType.descriptorIndex)
}
