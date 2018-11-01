package heap

import "jvmgo/classfile"

type Method struct {
	ClassMember
	maxLocals     uint
	maxStack      uint
	argsSlotCount uint
	code          []byte
}

func (method *Method) IsAbstract() bool {
	return HasFlag(method.accessFlags, ACC_ABSTRACT)
}

func (method *Method) IsStatic() bool {
	return HasFlag(method.accessFlags, ACC_STATIC)
}

func (method *Method) IsPublic() bool {
	return HasFlag(method.accessFlags, ACC_PUBLIC)
}

func (method *Method) IsProtected() bool {
	return HasFlag(method.accessFlags, ACC_PROTECTED)
}

func (method *Method) IsDefault() bool {
	return !HasFlag(method.accessFlags, ACC_PUBLIC, ACC_PROTECTED, ACC_PRIVATE)
}

func (method *Method) IsPrivate() bool {
	return HasFlag(method.accessFlags, ACC_PRIVATE)
}

func (method *Method) ArgsSlotCount() uint {
	return method.argsSlotCount
}

func (method *Method) Code() []byte {
	return method.code
}

func (method *Method) MaxStack() uint {
	return method.maxStack
}

func (method *Method) MaxLocals() uint {
	return method.maxLocals
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
	method.argsSlotCount = cnt
}

func (method *Method) IsAccessibleTo(otherClz *Class) bool {
	clz := method.class
	if clz == otherClz {
		return true
	}

	if method.IsPublic() {
		return true
	}

	if method.IsProtected() && (clz.SamePackage(otherClz) || otherClz.IsSubClassOf(clz)) {
		return true
	}

	if method.IsDefault() && clz.SamePackage(otherClz) {
		return true
	}

	return false
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
