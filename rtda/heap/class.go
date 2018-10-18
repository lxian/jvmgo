package heap

type Class struct {
	// CL
	classLoader *ClassLoader
	// basic info
	accessFlags uint16
	name string
	superClassName string
	superClass *Class
	interfaceNames string
	interfaces []*Class
	// constant pool
	constantPool ConstantPool
	fields []Field
	methods []Method
	// var
	instanceSlotCount uint
	staticSlotCount uint
	staticVars *Slots
}
