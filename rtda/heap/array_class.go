package heap

func (class *Class) NewArray(count uint) *Object {
	if !class.IsArray() {
		panic("Non-array class: " + class.name)
	}

	obj := &Object{}
	obj.class = class
	switch class.name {
	case ARR_BOOL  : obj.fields = make([]int8, count)
	case ARR_BYTE  :obj.fields = make([]int8, count)
	case ARR_CHAR  :obj.fields = make([]int16, count)
	case ARR_SHORT :obj.fields = make([]int16, count)
	case ARR_INT   :obj.fields = make([]int32, count)
	case ARR_LONG  :obj.fields = make([]int64, count)
	case ARR_FLOAT :obj.fields = make([]float32, count)
	case ARR_DOUBLE:obj.fields = make([]float64, count)
	default: obj.fields = make([]*Object, count)
	}
	return obj
}

func (class *Class) ArrayClass() *Class {
	arrClzName := getArrayClassName(class.name)
	return class.classLoader.loadArrayClass(arrClzName)
}

func (class *Class) IsArray() bool {
	return class.name[0] == '['
}
