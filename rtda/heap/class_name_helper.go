package heap

import "strings"

func getArrayClassName(className string) string {
	return "[" + getDescriptor(className)
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
