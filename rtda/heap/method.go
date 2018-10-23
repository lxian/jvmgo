package heap

import "jvmgo/classfile"

type Method struct {
	ClassMember
	maxLocals uint
	maxStack  uint
	code      []byte
}

func (m *Method) MaxStack() uint {
	return m.maxStack
}

func (m *Method) MaxLocals() uint {
	return m.maxLocals
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
	}

	return methods
}


