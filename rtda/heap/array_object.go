package heap

// fetch typed array
func (object *Object) Bytes() []int8 {
	return object.fields.([]int8)
}

func (object *Object) Shorts() []int16 {
	return object.fields.([]int16)
}

func (object *Object) Chars() []int16 {
	return object.fields.([]int16)
}

func (object *Object) Ints() []int32 {
	return object.fields.([]int32)
}

func (object *Object) Longs() []int64 {
	return object.fields.([]int64)
}

func (object *Object) Floats() []float32 {
	return object.fields.([]float32)
}

func (object *Object) Doubles() []float64 {
	return object.fields.([]float64)
}

func (object *Object) Refs() []*Object {
	return object.fields.([]*Object)
}

// length
func (object *Object) ArrayLength() int32 {
	switch object.fields.(type) {
	case []int8:
		return int32(len(object.fields.([]int8)))
	case []int16:
		return int32(len(object.fields.([]int16)))
	case []int32:
		return int32(len(object.fields.([]int32)))
	case []int64:
		return int32(len(object.fields.([]int64)))
	case []float32:
		return int32(len(object.fields.([]float32)))
	case []float64:
		return int32(len(object.fields.([]float64)))
	case []*Object:
		return int32(len(object.fields.([]*Object)))
	default:
		panic("Unrecognized Array type")
	}
}

func (object *Object) AssertArrIdx(idx int32) {
	if idx < 0 || idx >= object.ArrayLength() {
		panic("java.lang.ArrayIndexOutOfBoundsException")
	}
}
