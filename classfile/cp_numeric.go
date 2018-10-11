package classfile

import "math"

// Integer
type ConstantIntegerInfo struct {
	value int32
}

func (integerInfo *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	integerInfo.value = int32(reader.readUint32())
}

// Long
type ConstantLongInfo struct {
	value int64
}

func (longInfo *ConstantLongInfo) readInfo(reader *ClassReader) {
	longInfo.value = int64(reader.readUint64())
}

// Float
type ConstantFloatInfo struct {
	value float32
}

func (floatInfo *ConstantFloatInfo) readInfo(reader *ClassReader) {
	floatInfo.value = math.Float32frombits(reader.readUint32())
}

// Double
type ConstantDoubleInfo struct {
	value float64
}

func (doubleInfo *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	doubleInfo.value = math.Float64frombits(reader.readUint64())
}
