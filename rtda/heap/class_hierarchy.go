package heap

func (class *Class) IsSubClassOf(other *Class) bool {
	for c := class.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

func (class *Class) Implemented(other *Class) bool {
	if class == other {
		return true
	}
	for _, iface := range class.interfaces {
		if iface.IsSubClassOf(other) {
			return true
		}
	}
	if class.superClass != nil {
		return class.superClass.Implemented(other)
	}
	return false
}

func (class *Class) IsInterface() bool {
	return HasFlag(class.accessFlags, ACC_INTERFACE)
}

func (class *Class) IsJLObject() bool {
	return class.name == "java/lang/Object"
}

func (class *Class) IsJLConeable() bool {
	return class.name == "java/lang/Cloneable"
}

func (class *Class) IsJLSerializable() bool {
	return class.name == "java/io/Serializable"
}

func (class *Class) isAssignableFrom(other *Class) bool {
	t, s := class, other
	if s == t {
		return true
	}

	if !s.IsArray() {
		if !t.IsInterface() {
			return s.IsSubClassOf(t)
		} else {
			return s.Implemented(t)
		}
	}

	if s.IsInterface() {
		return t.IsJLObject() || s.isSubInterfaceOf(t)
	}

	if s.IsArray() {
		if !t.IsArray() {
			if !t.IsInterface() {
				return t.IsJLObject()
			} else {
				return t.IsJLSerializable() || t.IsJLConeable()
			}
		} else {
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || tc.isAssignableFrom(sc)
		}
	}

	return false
}
