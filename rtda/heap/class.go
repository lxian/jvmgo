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
	// class init
	initStarted bool
}

func (class *Class) ClassLoader() *ClassLoader {
	return class.classLoader
}

func (class *Class) FindStaticMethod(name string, desc string) *Method {
	for _, method := range class.methods {
		if method.IsStatic() && method.name == name && method.descriptor == desc {
			return method
		}
	}
	return nil
}

func (class *Class) SuperClass() *Class {
	return class.superClass
}

func (class *Class) FindInitMethod() *Method {
	return class.FindStaticMethod("<clinit>", "()V")
}

func (class *Class) StartInit() {
	class.initStarted = true
}

func (class *Class) InitStarted() bool {
	return class.initStarted
}

func (class *Class) Name() string {
	return class.name
}

func (class *Class) Methods() []*Method {
	return class.methods
}

func (class *Class) StaticVars() Slots {
	return class.staticVars
}

func (class *Class) AccessFlags() uint16 {
	return class.accessFlags
}

func (class *Class) ConstantPool() *ConstantPool {
	return class.constantPool
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

func (class *Class) SamePackage(other *Class) bool {
	return class.packageName() == other.packageName()
}

func (class *Class) IsAccessibleTo(other *Class) bool {
	return HasFlag(class.accessFlags, ACC_PUBLIC) || class.packageName() == other.packageName()
}
