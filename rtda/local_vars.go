package rtda

import "math"

type LocalVars []Slot

func newLocalVars(maxLocalVarSize uint) LocalVars {
	return make([]Slot, maxLocalVarSize)
}

// Int
func (locals LocalVars) SetInt(index uint, value int32) {
	locals[index].num = value
}

func (locals LocalVars) GetInt(index uint) int32 {
	return locals[index].num
}

// Float
func (locals LocalVars) SetFloat(index uint, value float32) {
	bits := math.Float32bits(value)
	locals[index].num = int32(bits)
}

func (locals LocalVars) GetFloat(index uint) float32 {
	return float32(math.Float32frombits(uint32(locals[index].num)))
}

// Long
func (locals LocalVars) SetLong(index uint, value int64) {
	locals[index].num = int32(value)         // lower bits
	locals[index+1].num = int32(value >> 32) // higher bits
}

func (locals LocalVars) GetLong(index uint) int64 {
	lower := locals[index].num
	higher := locals[index+1].num
	return int64(higher)<<32 | (int64(lower) & 0x00000000FFFFFFFF)
}

// Double
func (locals LocalVars) SetDouble(index uint, value float64) {
	bits := math.Float64bits(value)
	locals.SetLong(index, int64(bits))
}

func (locals LocalVars) GetDouble(index uint) float64 {
	return math.Float64frombits(uint64(locals.GetLong(index)))
}

// Ref
func (locals LocalVars) SetRef(index uint, ref *Object) {
	locals[index].ref = ref
}

func (locals LocalVars) GetRef(index uint) *Object {
	return locals[index].ref
}
