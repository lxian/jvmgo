package heap

import (
	"jvmgo/classfile"
	"strings"
)

type Class struct {
	// CL
	classLoader *ClassLoader
	// basic info
	accessFlags    uint16
	name           string
	superClassName string
	superClass     *Class
	interfaceNames []string
	interfaces     []*Class
	// constant pool
	constantPool *ConstantPool
	fields       []*Field
	methods      []*Method
	// var
	instanceSlotCount uint
	staticSlotCount   uint
	staticVars        Slots
}

func (c *Class) Methods() []*Method {
	return c.methods
}

func (c *Class) StaticVars() Slots {
	return c.staticVars
}

func (c *Class) AccessFlags() uint16 {
	return c.accessFlags
}

func (c *Class) ConstantPool() *ConstantPool {
	return c.constantPool
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

func (class *Class) packageName() string {
	if i := strings.LastIndex(class.name, "/"); i >= 0 {
		return class.name[:i]
	}
	return ""
}

func (class *Class) IsAccessibleTo(other *Class) bool {
	return HasFlag(class.accessFlags, ACC_PUBLIC) || class.packageName() == other.packageName()
}
