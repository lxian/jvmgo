package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

func (self *ClassReader) readUint8() uint8 {
	value := self.data[0]
	self.data = self.data[1:]
	return value
}

func (self *ClassReader) readUint16() uint16 {
	value := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return value
}

func (self *ClassReader) readUint32() uint32 {
	value := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return value
}

func (self *ClassReader) readUint64() uint64 {
	value := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return value
}

func (self *ClassReader) readUint16s() []uint16 {
	count := self.readUint16()
	values := make([]uint16, count)
	for i := range values {
		values[i] = self.readUint16()
	}
	return values
}

func (self *ClassReader) readBytes(count uint32) []byte {
	bytes := self.data[:count]
	self.data = self.data[count:]
	return bytes
}
