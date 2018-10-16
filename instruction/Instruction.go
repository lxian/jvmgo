package instruction

import "jvmgo/rtda"

type Instruction interface {
	FetchOperands(reader *ByteCodeReader)
	Execute(frame *rtda.Frame)
}

// NOP
type NoOperandsInstruction struct{}

func (nopInst *NoOperandsInstruction) FetchOperands(reader *ByteCodeReader) {
}

// Branch
type BranchInstruction struct {
	Offset int
}

func (inst *BranchInstruction) FetchOperands(reader *ByteCodeReader) {
	inst.Offset = int(reader.ReadInt16())
}

// Index 8
type Index8Instruction struct {
	Index uint
}

func (inst *Index8Instruction) FetchOperands(reader *ByteCodeReader) {
	inst.Index = uint(reader.ReadUint8())
}

// Index 16
type Index16Instruction struct {
	Index uint
}

func (inst *Index16Instruction) FetchOperands(reader *ByteCodeReader) {
	inst.Index = uint(reader.ReadUint16())
}
