package heap

import "jvmgo/classfile"

type Method struct {
	ClassMember
	maxLocals uint
	maxStack  uint
	argsCount uint
	code      []byte
}

func (m *Method) IsStatic() bool {
	return HasFlag(m.accessFlags, ACC_STATIC)
}

func (m *Method) ArgsCount() uint {
	return m.argsCount
}

func (m *Method) Code() []byte {
	return m.code
}

func (m *Method) MaxStack() uint {
	return m.maxStack
}

func (m *Method) MaxLocals() uint {
	return m.maxLocals
}

func (method *Method) calcArgsCount() {
	methodDescriptor := parseMethodDescriptors(method.descriptor)
	var cnt uint = 0
	for _, paramType := range methodDescriptor.paramTypes {
		switch paramType {
		case "J", "D":
			cnt += 2
		default:
			cnt += 1
		}
	}
	if !HasFlag(method.accessFlags, ACC_STATIC) {
		cnt += 1 // the implicit *this*
	}
	method.argsCount = cnt
}

func (method *Method) IsAccessibleTo(other *Class) bool {
	if !method.class.IsAccessibleTo(other) {
		return false
	}

}

func newMethods(class *Class, methodInfos []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(methodInfos))
	for i, info := range methodInfos {
		method := &Method{}
		method.copyInfoFromMemberInfo(info)
		method.class = class

		if code := info.FindCodeAttribute(); code != nil {
			method.maxLocals = uint(code.MaxLocals())
			method.maxStack = uint(code.MaxStack())
			method.code = code.Code()
		}

		methods[i] = method
		methods[i].calcArgsCount()
	}

	return methods
}

