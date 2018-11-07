package heap

// fetch typed array
func (object *Object) Bytes() []int8 {
	return object.vars.([]int8)
}

func (object *Object) Shorts() []int16 {
	return object.vars.([]int16)
}

func (object *Object) Chars() []uint16 {
	return object.vars.([]uint16)
}

func (object *Object) Ints() []int32 {
	return object.vars.([]int32)
}

func (object *Object) Longs() []int64 {
	return object.vars.([]int64)
}

func (object *Object) Floats() []float32 {
	return object.vars.([]float32)
}

func (object *Object) Doubles() []float64 {
	return object.vars.([]float64)
}

func (object *Object) Refs() []*Object {
	return object.vars.([]*Object)
}

// length
func (object *Object) ArrayLength() int32 {
	switch object.vars.(type) {
	case []int8:
		return int32(len(object.vars.([]int8)))
	case []int16:
		return int32(len(object.vars.([]int16)))
	case []int32:
		return int32(len(object.vars.([]int32)))
	case []int64:
		return int32(len(object.vars.([]int64)))
	case []float32:
		return int32(len(object.vars.([]float32)))
	case []float64:
		return int32(len(object.vars.([]float64)))
	case []*Object:
		return int32(len(object.vars.([]*Object)))
	default:
		panic("Unrecognized Array type")
	}
}

func (object *Object) AssertArrIdx(idx int32) {
	if idx < 0 || idx >= object.ArrayLength() {
		panic("java.lang.ArrayIndexOutOfBoundsException")
	}
}
