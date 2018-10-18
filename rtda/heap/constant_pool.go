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

func newConstant(constantInfo classfile.ConstantInfo) Constant {
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

	}
	panic(fmt.Sprintf("Unrecognized constant type %v", constantInfo))
}

func (cp *ConstantPool) GetConstant(index uint) Constant {
	if c := cp.constants[index]; c != nil {
		return c
	}
	panic(fmt.Sprintf("Empty constant at %d", index))
}

