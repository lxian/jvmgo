package heap

func (class *Class) isSubClassOf(other *Class) bool {
	if class == other {
		return true
	}

	if class.superClass != nil {
		return class.superClass.isSubClassOf(other)
	}
	return false
}

func (class *Class) Implemented(other *Class) bool {
	if class == other {
		return true
	}
	for _, iface := range class.interfaces {
		if iface.isSubClassOf(other) {
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

func (class *Class) isAssignableFrom(other *Class) bool {
	if other.IsInterface() {
		if class.IsInterface() {
			return other.Implemented(class)
		} else {
			return class.name == "java/lang/Object"
		}
	} else {
		if class.IsInterface() {
			return other.Implemented(class)
		} else {
			return other.isSubClassOf(class)
		}
	}
}
