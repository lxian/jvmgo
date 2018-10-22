package heap

import "math"

type Slot struct {
	num int32
	ref *Object
}

type Slots []Slot

func newSlots(maxLocalVarSize uint) Slots {
	return make([]Slot, maxLocalVarSize)
}

// Int
func (slots Slots) SetInt(index uint, value int32) {
	slots[index].num = value
}

func (slots Slots) GetInt(index uint) int32 {
	return slots[index].num
}

// Float
func (slots Slots) SetFloat(index uint, value float32) {
	bits := math.Float32bits(value)
	slots[index].num = int32(bits)
}

func (slots Slots) GetFloat(index uint) float32 {
	return float32(math.Float32frombits(uint32(slots[index].num)))
}

// Long
func (slots Slots) SetLong(index uint, value int64) {
	slots[index].num = int32(value)         // lower bits
	slots[index+1].num = int32(value >> 32) // higher bits
}

func (slots Slots) GetLong(index uint) int64 {
	lower := slots[index].num
	higher := slots[index+1].num
	return int64(higher)<<32 | (int64(lower) & 0x00000000FFFFFFFF)
}

// Double
func (slots Slots) SetDouble(index uint, value float64) {
	bits := math.Float64bits(value)
	slots.SetLong(index, int64(bits))
}

func (slots Slots) GetDouble(index uint) float64 {
	return math.Float64frombits(uint64(slots.GetLong(index)))
}

// Ref
func (slots Slots) SetRef(index uint, ref *Object) {
	slots[index].ref = ref
}

func (slots Slots) GetRef(index uint) *Object {
	return slots[index].ref
}

