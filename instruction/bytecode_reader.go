package instruction

type ByteCodeReader struct {
	code []byte
	pc int
}

func (reader *ByteCodeReader) Reset(code []byte, pc int) {
	reader.code = code
	reader.pc = pc
}

func (reader *ByteCodeReader) ReadUint8() uint8 {
	b := reader.code[reader.pc]
	reader.pc += 1
	return b
}

func (reader *ByteCodeReader) ReadInt8() int8 {
	return int8(reader.ReadUint8())
}

func (reader *ByteCodeReader) ReadUint16() uint16 {
	return uint16(reader.ReadUint8()) << 8 | (uint16(reader.ReadUint8()) & 0x00FF)
}

func (reader *ByteCodeReader) ReadInt16() int16 {
	return int16(reader.ReadUint16())
}

func (reader *ByteCodeReader) ReadUint32() uint32 {
	return uint32(reader.ReadUint16()) << 16  | (uint32(reader.ReadUint16()) & 0x0000FFFF)
}

func (reader *ByteCodeReader) ReadInt32() int32 {
	return int32(reader.ReadUint32())
}

func (reader *ByteCodeReader) ReadInt32s() {
}

func (reader *ByteCodeReader) SkipPadding() {
}


