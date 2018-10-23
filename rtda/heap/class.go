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

func (class *Class) isAccessibleTo(other *Class) bool {
	return HasFlag(class.accessFlags, ACC_PUBLIC) || class.packageName() == other.packageName()
}

func (class *Class) isSubClassOf(other *Class) bool {
	if class == other {
		return true
	}

	if class.superClass != nil {
		return class.superClass.isSubClassOf(other)
	}
	return false
}

func (class *Class) Implemented(other *Class) bool {
	if class == other {
		return true
	}
	for _, iface := range class.interfaces {
		if iface.isSubClassOf(other) {
			return true
		}
	}
	if class.superClass != nil {
		return class.superClass.Implemented(other)
	}
	return false
}

func (class *Class) IsInterface() bool {
	return HasFlag(class.accessFlags, ACC_INTERFACE)
}

func (class *Class) isAssignableFrom(other *Class) bool {
	if other.IsInterface() {
		if class.IsInterface() {
			return other.Implemented(class)
		} else {
			return class.name == "java/lang/Object"
		}
	} else {
		if class.IsInterface() {
			return other.Implemented(class)
		} else {
			return other.isSubClassOf(class)
		}
	}
}
