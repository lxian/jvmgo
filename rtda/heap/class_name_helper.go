package heap

import "strings"

func getArrayClassName(className string) string {
	return "[" + getDescriptor(className)
}

func getComponentClassName(className string) string {
	return getClassName(className[1:])
}

func getClassName(descriptor string) string {
	if descriptor[0] == '[' {
		return descriptor
	}

	for clz, desc := range primitives_mapping {
		if string(desc) == descriptor {
			return clz
		}
	}

	if descriptor[0] == 'L' {
		return descriptor[1 : len(descriptor)-1]
	}
	return descriptor
}

func getDescriptor(className string) string {
	// array type
	if className[0] == '[' {
		return className
	}

	// primitive type
	if desc, ok := getPrimitiveTypes(className); ok {
		return desc
	}

	// object
	return "L" + className + ";"
}

func getPrimitiveTypes(className string) (string, bool) {
	if desc, ok := primitives_mapping[className]; ok {
		return string(desc), true
	} else {
		return "", false
	}
}

func GetJavaClassName(className string) string {
	return strings.Replace(className, "/", ".", -1)
}
