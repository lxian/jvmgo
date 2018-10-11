package classfile

const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

func readConstantInfo(reader *ClassReader, constantPool ConstantPool) ConstantInfo {
	flag := reader.readUint8()
	constantInfo := newConstantInfo(flag, constantPool)
	constantInfo.readInfo(reader)
	return constantInfo
}

func newConstantInfo(flag uint8, constantPool ConstantPool) ConstantInfo {
	switch flag {
	case CONSTANT_Class:
		return &ConstantClassInfo{constantPool: constantPool}
	case CONSTANT_Fieldref, CONSTANT_Methodref, CONSTANT_InterfaceMethodref:
		return &ConstantMemberRefInfo{constantPool: constantPool}
	case CONSTANT_String:
		return &ConstantStringInfo{constantPool: constantPool}
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_NameAndType:
		return &ConstantNameAndType{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	}
	return nil
}
