package jump

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type GOTO struct { instruction.Index16Instruction }

func (inst *GOTO) Execute(frame *rtda.Frame) {
	instruction.Branch(frame, int(inst.Index))
}

type GOTOW struct { index uint32 }

func (inst *GOTOW) FetchOperands(reader *instruction.ByteCodeReader) {
	b1 := reader.ReadUint8()
	b2 := reader.ReadUint8()
	b3 := reader.ReadUint8()
	b4 := reader.ReadUint8()
	inst.index = uint32(b1) << 24 | uint32(b2) << 16 | uint32(b3) << 8 | uint32(b4)
}

func (inst *GOTOW) Execute(frame *rtda.Frame) {
	instruction.Branch(frame, int(inst.index))
}
