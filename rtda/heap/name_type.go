package heap

import "jvmgo/classfile"

type NameAndType struct {
	pool          *ConstantPool
	nameIdx       uint16
	descriptorIdx uint16
}

func (nameAndType *NameAndType) Name() string {
	return nameAndType.pool.GetConstant(uint(nameAndType.nameIdx)).(string)
}

func (nameAndType *NameAndType) Descriptor() string {
	return nameAndType.pool.GetConstant(uint(nameAndType.descriptorIdx)).(string)
}

func newNameAndType(pool *ConstantPool, info *classfile.ConstantNameAndType) *NameAndType {
	return &NameAndType{pool: pool, nameIdx: info.NameIndex(), descriptorIdx: info.DescriptorIndex()}
}
