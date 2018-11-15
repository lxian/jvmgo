package heap

import (
	"fmt"
	"jvmgo/classfile"
	"jvmgo/classpath"
)

type ClassLoader struct {
	classPath    *classpath.Classpath
	classMap     map[string]*Class
	verboseClass bool
}

func NewClassLaoder(cp *classpath.Classpath, verboseClass bool) *ClassLoader {
	loader := &ClassLoader{classPath: cp, classMap: make(map[string]*Class), verboseClass: verboseClass}
	loader.LoadBasicClasses()
	loader.loadPrimitiveClasses()
	return loader
}

func (loader *ClassLoader) loadPrimitiveClasses() {
	for primitiveType := range primitives_mapping {
		loader.loadPrimitiveClass(primitiveType)
	}
}

func (loader *ClassLoader) loadPrimitiveClass(className string) {
	class := &Class{
		classLoader: loader,
		accessFlags: ACC_PUBLIC,
		name:        className,
		initStarted: true,
	}
	class.jClassObj = loader.classMap["java/lang/Class"].NewObject()
	class.jClassObj.extra = class
	loader.classMap[className] = class
}

func (loader *ClassLoader) LoadBasicClasses() {
	jClzClz := loader.LoadClass("java/lang/Class")
	for _, clz := range loader.classMap {
		if clz.jClassObj == nil {
			clz.jClassObj = jClzClz.NewObject()
			clz.jClassObj.extra = clz
		}
	}
}

func (loader *ClassLoader) LoadClass(className string) *Class {
	if class, ok := loader.classMap[className]; ok {
		return class
	}

	var class *Class
	if className[0] == '[' {
		class = loader.loadArrayClass(className)
	} else {
		class = loader.loadNonArrayClass(className)
	}

	if jClzClz, ok := loader.classMap["java/lang/Class"]; ok {
		class.jClassObj = jClzClz.NewObject()
		class.jClassObj.extra = class
	}

	loader.classMap[class.name] = class
	return class
}

func (loader *ClassLoader) loadNonArrayClass(className string) *Class {
	data, entry := loader.readClass(className)
	class := loader.defineClass(data)
	link(class)

	if loader.verboseClass {
		fmt.Printf("[Load %s from %s]\n", className, entry)
	}
	return class
}

func (loader *ClassLoader) readClass(className string) ([]byte, classpath.Entry) {
	data, entry, err := loader.classPath.ReadClass(className)
	if err != nil {
		panic(fmt.Sprintf("Failed reading class %s, %v", className, err))
	}
	return data, entry
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
	link(class)

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

func (loader *ClassLoader) loadArrayClass(name string) *Class {
	return &Class{
		classLoader: loader,
		accessFlags: ACC_PUBLIC, // TODO
		name:        name,
		initStarted: true,
		superClass:  loader.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			loader.LoadClass("java/lang/Cloneable"),
			loader.LoadClass("java/io/Serializable"),
		},
	}
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
	cpIdx := uint(field.constantValueIndex)

	if cpIdx > 0 {
		constVal := class.constantPool.GetConstant(cpIdx)
		switch field.descriptor {
		case "B", "C", "I", "Z", "S":
			class.staticVars.SetInt(field.slotId, constVal.(int32))
		case "F":
			class.staticVars.SetFloat(field.slotId, constVal.(float32))
		case "D":
			class.staticVars.SetDouble(field.slotId, constVal.(float64))
		case "J":
			class.staticVars.SetLong(field.slotId, constVal.(int64))
		case "Ljava/lang/String;":
			class.staticVars.SetRef(field.slotId, InternedJString(constVal.(string), class.classLoader))
		default:
			panic(fmt.Sprintf("Field %v Constant Value: Bad descriptor", field))
		}
	}
}
