package native

var registry = map[string]NativeMethod{}

func key(className, methodName, methodDescriptor string) string {
	return className + "." + methodName + "." + methodDescriptor
}

func RegisterNativeMethod(className, methodName, methodDescriptor string, nativeMethod NativeMethod) {
	key := key(className, methodName, methodDescriptor)
	registry[key] = nativeMethod
}

func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	key := key(className, methodName, methodDescriptor)
	if nm, ok := registry[key]; ok {
		return nm
	}
	//check and return empty registerNatives func
	if methodName == "registerNatives" && methodDescriptor == "()V" {
		return EmptyNativeMethod
	}

	panic("java.lang.UnsatisfiedLinkError: " + key)
}
