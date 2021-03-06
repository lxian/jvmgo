package rtda

import (
	"jvmgo/rtda/heap"
	"math"
)

type OperandStack struct {
	size     uint
	operands []Slot
}

func newOperandStack(maxOperandStackSize uint) *OperandStack {
	return &OperandStack{size: 0, operands: make([]Slot, maxOperandStackSize)}
}

func (stack *OperandStack) Clear() {
	stack.size = 0
	for _, slot := range stack.operands {
		slot.ref = nil
	}
}

// Slot
func (stack *OperandStack) PeekSlot() Slot {
	return stack.operands[stack.size-1]
}

func (stack *OperandStack) PeekSlotBelow(cnt uint) Slot {
	topIdx := stack.size - 1
	return stack.operands[topIdx-cnt]
}

func (stack *OperandStack) PushSlot(slot Slot) {
	stack.operands[stack.size] = slot
	stack.size += 1
}

func (stack *OperandStack) PopSlot() Slot {
	stack.size -= 1
	return stack.operands[stack.size]
}

// Int
func (stack *OperandStack) PushInt(value int32) {
	stack.operands[stack.size].num = value
	stack.size += 1
}

func (stack *OperandStack) PopInt() int32 {
	stack.size -= 1
	return stack.operands[stack.size].num
}

// Float
func (stack *OperandStack) PushFloat(value float32) {
	stack.operands[stack.size].num = int32(math.Float32bits(value))
	stack.size += 1
}

func (stack *OperandStack) PopFloat() float32 {
	stack.size -= 1
	return math.Float32frombits(uint32(stack.operands[stack.size].num))
}

// Long
func (stack *OperandStack) PushLong(value int64) {
	stack.PushInt(int32(value))       // lower
	stack.PushInt(int32(value >> 32)) // higher
}

func (stack *OperandStack) PopLong() int64 {
	higher := stack.PopInt()
	lower := stack.PopInt()
	return int64(higher)<<32 | (int64(lower) & 0x00000000FFFFFFFF)
}

// Double
func (stack *OperandStack) PushDouble(value float64) {
	stack.PushLong(int64(math.Float64bits(value)))
}

func (stack *OperandStack) PopDouble() float64 {
	return math.Float64frombits(uint64(stack.PopLong()))
}

// Ref
func (stack *OperandStack) PushRef(ref *heap.Object) {
	stack.operands[stack.size].ref = ref
	stack.size += 1
}

func (stack *OperandStack) PopRef() *heap.Object {
	stack.size -= 1
	val := stack.operands[stack.size].ref
	stack.operands[stack.size].ref = nil
	return val
}

func (stack *OperandStack) PushBoolean(val bool) {
	if val {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
func (stack *OperandStack) PopBoolean() bool {
	return stack.PopInt() == 1
}
