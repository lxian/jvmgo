package heap

func (class *Class) NewArray(count uint) *Object {
	if !class.IsArray() {
		panic("Non-array class: " + class.name)
	}

	obj := &Object{}
	obj.class = class
	switch class.name {
	case ARR_BOOL:
		obj.vars = make([]int8, count)
	case ARR_BYTE:
		obj.vars = make([]int8, count)
	case ARR_CHAR:
		obj.vars = make([]uint16, count)
	case ARR_SHORT:
		obj.vars = make([]int16, count)
	case ARR_INT:
		obj.vars = make([]int32, count)
	case ARR_LONG:
		obj.vars = make([]int64, count)
	case ARR_FLOAT:
		obj.vars = make([]float32, count)
	case ARR_DOUBLE:
		obj.vars = make([]float64, count)
	default:
		obj.vars = make([]*Object, count)
	}
	return obj
}

func (class *Class) ArrayClass() *Class {
	arrClzName := getArrayClassName(class.name)
	return class.classLoader.LoadClass(arrClzName)
}

func (class *Class) ComponentClass() *Class {
	if !class.IsArray() {
		return nil
	}
	compClz := getComponentClassName(class.name)
	return class.classLoader.LoadClass(compClz)
}

func (class *Class) IsArray() bool {
	return class.name[0] == '['
}

func (class *Class) IsPrimitive() bool {
	_, ok := primitives_mapping[class.name]
	return ok
}
