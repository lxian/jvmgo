package heap

import "jvmgo/classfile"

type Method struct {
	ClassMember
	maxLocals uint
	maxStack  uint
	argsCount uint
	code      []byte
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

func (method *Method) calcArgsCount() uint {
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
		methods[i].argsCount = methods[i].calcArgsCount()
	}

	return methods
}
