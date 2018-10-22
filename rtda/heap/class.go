package heap

import "jvmgo/classfile"

type Class struct {
	// CL
	classLoader *ClassLoader
	// basic info
	accessFlags uint16
	name string
	superClassName string
	superClass *Class
	interfaceNames []string
	interfaces []*Class
	// constant pool
	constantPool *ConstantPool
	fields []*Field
	methods []*Method
	// var
	instanceSlotCount uint
	staticSlotCount uint
	staticVars Slots
}

func newClass(classfile *classfile.ClassFile) *Class {
	cls := &Class{}
	// basic info
	cls.accessFlags = classfile.AccessFlags()
	cls.name = classfile.ClassName()
	cls.superClassName = classfile.SuperClassName()
	cls.interfaceNames = classfile.InterfaceNames()
	// const pool
	cls.constantPool = newConstantPool(cls, classfile.ConstantPool())
	cls.fields = newFields(cls, classfile.Field())
	cls.methods = newMethods(cls, classfile.Methods())

	return cls
}

func (class *Class) packageName {
	return class.name.
}

func (class *Class) isAccessibleTo(target *Class) {
	class.name.
}

