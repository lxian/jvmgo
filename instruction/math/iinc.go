package math

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type IINC struct {
	index uint8
	constValue int32
}

func (inst *IINC) FetchOperands(reader *instruction.ByteCodeReader) {
	inst.index = reader.ReadUint8()
	inst.constValue = int32(reader.ReadInt8())
}

func (inst *IINC) Execute(frame *rtda.Frame) {
	v := frame.LocalVars().GetInt(uint(inst.index))
	v += inst.constValue
	frame.LocalVars().SetInt(uint(inst.index), v)
}
