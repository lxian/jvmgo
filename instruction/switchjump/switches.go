package switchjump

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type TABLE_SWITCH struct {
	defaultOffset int32
	lowOffset     int32
	highOffset    int32
	table         []int32
}

func (inst *TABLE_SWITCH) FetchOperands(reader *instruction.ByteCodeReader) {
	reader.SkipPadding()
	inst.defaultOffset = reader.ReadInt32()
	inst.lowOffset = reader.ReadInt32()
	inst.highOffset = reader.ReadInt32()
	inst.table = make([]int32, inst.highOffset-inst.lowOffset+1)
	for i := range inst.table {
		inst.table[i] = reader.ReadInt32()
	}
}

func (inst *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()
	if index < inst.lowOffset || index > inst.highOffset {
		instruction.Branch(frame, int(inst.defaultOffset))
	} else {
		instruction.Branch(frame, int(inst.table[index-inst.lowOffset]))
	}
}

type LOOKUP_SWITCH struct {
	defaultOffset int32
	npairs        int32
	keys          []int32
	offsets       []int32
}

func (inst *LOOKUP_SWITCH) FetchOperands(reader *instruction.ByteCodeReader) {
	reader.SkipPadding()
	inst.defaultOffset = reader.ReadInt32()
	inst.npairs = reader.ReadInt32()
	inst.keys = make([]int32, inst.npairs)
	inst.offsets = make([]int32, inst.npairs)
	for i := range inst.keys {
		inst.keys[i] = reader.ReadInt32()
		inst.offsets[i] = reader.ReadInt32()
	}
}

func (inst *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	offset := inst.defaultOffset
	for i, k := range inst.keys {
		if k == key {
			offset = inst.offsets[i]
			break
		}
	}

	instruction.Branch(frame, int(offset))
}
