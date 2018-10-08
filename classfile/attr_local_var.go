package classfile

type LocalVariableEntry struct {
	startPc         uint16
	length          uint16
	nameIndex       uint16
	descriptorIndex uint16
	index           uint16
}

type LocalVariableTableAttribute struct {
	entries []LocalVariableEntry
}

func (attr *LocalVariableTableAttribute) readInfo(reader *ClassReader) {
	entryCount := uint16(reader.readUint16())
	entries := make([]LocalVariableEntry, entryCount)
	for i := range entries {
		entries[i] = LocalVariableEntry{
			startPc:         reader.readUint16(),
			length:          reader.readUint16(),
			nameIndex:       reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			index:           reader.readUint16(),
		}
	}
	attr.entries = entries
}
