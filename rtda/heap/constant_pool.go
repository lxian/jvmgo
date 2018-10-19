package heap

import (
	"jvmgo/classfile"
	"fmt"
)

type Constant interface {}

type ConstantPool struct {
	class *Class
	constants []Constant
}

func newConstantPool(class *Class, cfConstantPool classfile.ConstantPool) *ConstantPool {
	constants := make([]Constant, len(cfConstantPool))
	for i := 1; i < len(cfConstantPool) ;{
		info := cfConstantPool[i]
		constants[i] = newConstant(info)
		switch info.(type) {
		case *classfile.ConstantLongInfo, classfile.ConstantDoubleInfo:
			i += 2
		default:
			i += 1
		}
	}

	return &ConstantPool{constants:constants, class:class}
}

func newConstant(cp *ConstantPool, constantInfo classfile.ConstantInfo) Constant {
	switch constantInfo.(type) {
	// numeric
	case *classfile.ConstantIntegerInfo:
		return constantInfo.(*classfile.ConstantIntegerInfo).Value()
	case *classfile.ConstantLongInfo:
		return constantInfo.(*classfile.ConstantLongInfo).Value()
	case *classfile.ConstantFloatInfo:
		return constantInfo.(*classfile.ConstantFloatInfo).Value()
	case *classfile.ConstantDoubleInfo:
		return constantInfo.(*classfile.ConstantDoubleInfo).Value()
	// string
	case *classfile.ConstantUtf8Info:
		return constantInfo.(*classfile.ConstantUtf8Info).Value()
	case *classfile.ConstantStringInfo:
		return constantInfo.(*classfile.ConstantStringInfo).String()
	// sym ref
	case *classfile.ConstantClassInfo:
		return newClassRef(cp, constantInfo.(*classfile.ConstantClassInfo))
	case *classfile.ConstantFieldRefInfo:
		return newFieldRef(cp, constantInfo.(*classfile.ConstantFieldRefInfo))
	case *classfile.ConstantMethodRefInfo:
		return newMethodRef(cp, constantInfo.(*classfile.ConstantMethodRefInfo))
	case *classfile.ConstantInterfaceMethodRefInfo:
		return newInterfaceMethodRef(cp, constantInfo.(*classfile.ConstantInterfaceMethodRefInfo))
	}
	panic(fmt.Sprintf("Unrecognized constant type %v", constantInfo))
}

func (cp *ConstantPool) GetConstant(index uint) Constant {
	if c := cp.constants[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("Empty constant at %d", index))
}

