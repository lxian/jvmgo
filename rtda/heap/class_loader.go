package heap

import (
	"fmt"
	"jvmgo/classfile"
	"jvmgo/classpath"
)

type ClassLoader struct {
	classPath *classpath.Classpath
	classMap  map[string]*Class
}

func newClassLaoder(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{classPath: cp, classMap: make(map[string]*Class)}
}

func (loader *ClassLoader) LoadClass(className string) *Class {
	if loader.classMap[className] != nil {
		return loader.classMap[className]
	}

	class := loader.loadNonArrayClass(className)
	loader.classMap[class.name] = class
	return class
}

func (loader *ClassLoader) loadNonArrayClass(className string) *Class {
	data := loader.readClass(className)
	class := loader.defineClass(data)
	link(class)
	return loader.defineClass(loader.readClass(className))
}

func (loader *ClassLoader) readClass(className string) []byte {
	data, _, err := loader.classPath.ReadClass(className)
	if err != nil {
		panic(fmt.Sprintf("Failed reading class %s, %v", className, err))
	}
	return data
}

func (loader *ClassLoader) defineClass(codebyte []byte) *Class {
	classFile, err := classfile.ParseClassBytes(codebyte)
	if err != nil {
		panic(err)
	}

	class := newClass(classFile)
	class.classLoader = loader
	class.superClass = loader.resolveSuperClass(class)
	class.interfaces = loader.resolveInterfaces(class)

	return class
}

func (loader *ClassLoader) resolveSuperClass(class *Class) *Class {
	if class.name == "java/lang/Object" {
		return nil
	}
	return loader.LoadClass(class.superClassName)
}

func (loader *ClassLoader) resolveInterfaces(class *Class) []*Class {
	interfaces := make([]*Class, len(class.interfaceNames))
	for i, name := range class.interfaceNames {
		interfaces[i] = loader.LoadClass(name)
	}
	return interfaces
}

func link(class *Class) {
	validate(class)
	prepare(class)
}

func validate(class *Class) {
	// DO NOTHING
}

func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

func isLongOrDouble(field *Field) bool {
	return field.descriptor == "D" /* double */ || field.descriptor == "J" /* long */
}

func calcStaticFieldSlotIds(class *Class) {
	fields := class.fields
	var idx uint = 0
	for _, field := range fields {
		if !HasFlag(field.accessFlags, ACC_STATIC) {
			continue
		}

		field.slotId = idx
		if isLongOrDouble(field) {
			idx += 2
		} else {
			idx++
		}
	}
	class.staticSlotCount = idx
}

func calcInstanceFieldSlotIds(class *Class) {
	var idx uint = 0
	if class.superClass != nil {
		idx = class.superClass.instanceSlotCount
	}
	fields := class.fields
	for _, field := range fields {
		if HasFlag(field.accessFlags, ACC_STATIC) {
			continue
		}

		field.slotId = idx
		if isLongOrDouble(field) {
			idx += 2
		} else {
			idx++
		}
	}
	class.instanceSlotCount = idx
}

func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if HasFlag(field.accessFlags, ACC_STATIC, ACC_FINAL) {
			initStaticVar(class, field)
		}
	}
}

func initStaticVar(class *Class, field *Field) {
	idx := uint(field.constantValueIndex)
	constVal := class.constantPool.GetConstant(idx)

	switch field.descriptor {
	case "B", "C", "I", "Z", "S":
		class.staticVars.SetInt(field.slotId, constVal.(int32))
	case "F":
		class.staticVars.SetFloat(field.slotId, constVal.(float32))
	case "D":
		class.staticVars.SetDouble(field.slotId, constVal.(float64))
	case "L":
		class.staticVars.SetLong(field.slotId, constVal.(int64))
	default:
		panic(fmt.Sprintf("Field %v Constant Value: Bad descriptor", field))
	}
}
