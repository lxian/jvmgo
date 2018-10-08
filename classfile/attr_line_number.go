package classfile

type LineNumberEntry struct {
	startPc    uint16
	lineNumber uint16
}

type LineNumberTableAttribute struct {
	entries []LineNumberEntry
}

func (attr *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	entryCount := uint16(reader.readUint16())
	entries := make([]LineNumberEntry, entryCount)
	for i := range entries {
		entries[i] = LineNumberEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
	attr.entries = entries
}
