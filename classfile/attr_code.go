package classfile

type ExceptionEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

type CodeAttribute struct {
	constantPool   ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []ExceptionEntry
	attributes     []AttributeInfo
}

func (attr *CodeAttribute) readInfo(reader *ClassReader) {
	attr.maxStack = reader.readUint16()
	attr.maxLocals = reader.readUint16()

	codeLen := uint32(reader.readUint32())
	attr.code = reader.readBytes(codeLen)

	expLen := reader.readUint16()
	expTable := make([]ExceptionEntry, expLen)
	for i := range expTable {
		expTable[i] = ExceptionEntry{
			reader.readUint16(),
			reader.readUint16(),
			reader.readUint16(),
			reader.readUint16(),
		}
	}
	attr.exceptionTable = expTable

	attr.attributes = readAttributes(reader, attr.constantPool)
}

func (attr *CodeAttribute) MaxLocals() uint16 {
	return attr.maxLocals
}

func (attr *CodeAttribute) MaxStack() uint16 {
	return attr.maxStack
}

func (attr *CodeAttribute) Code() []byte {
	return attr.code
}
