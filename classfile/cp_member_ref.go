package classfile

type ConstantMemberRefInfo struct {
	constantPool	ConstantPool
	classIndex		uint16
	nameAndTypeIndex	uint16
}
type ConstantFieldRefInfo struct { ConstantMemberRefInfo }
type ConstantMethodRefInfo struct { ConstantMemberRefInfo }
type ConstantInterfaceRefInfo struct { ConstantMemberRefInfo }

func (memberRefInfo *ConstantMemberRefInfo) readInfo(reader *ClassReader) {
	memberRefInfo.classIndex = reader.readUint16()
	memberRefInfo.nameAndTypeIndex = reader.readUint16()
}

func (memberRefInfo *ConstantMemberRefInfo) ClassName() string {
	return memberRefInfo.constantPool.getClassName(memberRefInfo.classIndex)
}

func (memberRefInfo *ConstantMemberRefInfo) NameAndType() (string, string) {
	return memberRefInfo.constantPool.getNameAndType(memberRefInfo.nameAndTypeIndex)
}

