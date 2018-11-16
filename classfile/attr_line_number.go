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

func (table *LineNumberTableAttribute) GetLineNumber(pc int) int {
	for _, e := range table.entries {
		if int(e.startPc) == pc {
			return int(e.lineNumber)
		}
	}
	return -1
}
