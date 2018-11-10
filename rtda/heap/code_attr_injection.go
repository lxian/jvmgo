package heap

const (
	NATIVE_TRAP = 0xfe
)

func (method *Method) InjectNativeCodeAttr(returnType string) {
	method.maxStack = 4
	method.maxLocals = method.argsSlotCount
	switch returnType[0] {
	case VOID:
		method.code = []byte{NATIVE_TRAP, 0xb1}
	case DOUBLE:
		method.code = []byte{NATIVE_TRAP, 0xaf}
	case FLOAT:
		method.code = []byte{NATIVE_TRAP, 0xae}
	case LONG:
		method.code = []byte{NATIVE_TRAP, 0xad}
	case OBJECT, ARRAY:
		method.code = []byte{NATIVE_TRAP, 0xb0}
	default:
		method.code = []byte{NATIVE_TRAP, 0xac}
	}
}
