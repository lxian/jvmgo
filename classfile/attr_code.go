package classfile

import "golang.org/x/tools/go/analysis/passes/pkgfact/testdata/src/c"

type ExceptionEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (e *ExceptionEntry) CatchType() uint16 {
	return e.catchType
}

func (e *ExceptionEntry) HandlerPc() uint16 {
	return e.handlerPc
}

func (e *ExceptionEntry) EndPc() uint16 {
	return e.endPc
}

func (e *ExceptionEntry) StartPc() uint16 {
	return e.startPc
}

type CodeAttribute struct {
	constantPool   ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []ExceptionEntry
	attributes     []AttributeInfo
}

func (attr *CodeAttribute) Attributes() []AttributeInfo {
	return attr.attributes
}

func (attr *CodeAttribute) ExceptionTable() []ExceptionEntry {
	return attr.exceptionTable
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
