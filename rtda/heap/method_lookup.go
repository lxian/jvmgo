package heap

func lookupMethod(class *Class, name string, desc string) *Method {
	method := LookupMethodInClass(class, name, desc)
	if method == nil {
		method = LookupMethodInInterfaces(class.interfaces, name, desc)
	}
	return method
}

func lookupInterfaceMethod(iface *Class, name string, desc string) *Method {
	return LookupMethodInInterfaces([]*Class{iface}, name, desc)
}

func LookupMethodInClass(class *Class, name string, desc string) *Method {
	for clz := class; clz != nil; clz = clz.superClass {
		for _, method := range clz.methods {
			if method.name == name && method.descriptor == desc {
				return method
			}
		}
	}
	return nil
}

func LookupMethodInInterfaces(ifaces []*Class, name string, desc string) *Method {
	for _, iface := range ifaces {
		for _, method := range iface.methods {
			if method.name == name && method.descriptor == desc {
				return method
			}
		}

		method := LookupMethodInInterfaces(iface.interfaces, name, desc)
		if method != nil {
			return method
		}
	}
	return nil
}
