package classfile

// Method handle
type ConstantMethodHandleInfo struct {
	referenceKind uint8
	referenceIndex uint16
}

func (mh *ConstantMethodHandleInfo)readInfo(reader *ClassReader) {
	mh.referenceKind = reader.readUint8()
	mh.referenceIndex = reader.readUint16()
}

// Method Type
type ConstantMethodTypeInfo struct {
	descriptorIndex uint16
}

func (mt *ConstantMethodTypeInfo)readInfo(reader *ClassReader) {
	mt.descriptorIndex = reader.readUint16()
}

// Invoke Dynamic
type ConstantInvokeDynamicInfo struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex uint16
}

func (vd *ConstantInvokeDynamicInfo)readInfo(reader *ClassReader) {
	vd.bootstrapMethodAttrIndex = reader.readUint16()
	vd.nameAndTypeIndex = reader.readUint16()
}
