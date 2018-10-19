package heap

import (
	"jvmgo/classpath"
	"fmt"
	"jvmgo/classfile"
)

type ClassLoader struct {
	classPath *classpath.Classpath
	classMap  map[string]*Class
}

func newClassLaoder(cp *classpath.Classpath) *ClassLoader {
	return &ClassLoader{classPath:cp, classMap:make(map[string]*Class)}
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
	data, _, err := loader.classPath.ReadClass(className)
	if err != nil {
		panic(fmt.Sprintf("Failed reading class %s, %v", className, err))
	}
	return loader.defineClass(data)
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


