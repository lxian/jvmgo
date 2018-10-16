package constant

import (
	"jvmgo/instruction"
	"jvmgo/rtda"
)

// byte sign-extended to int
type BIPUSH struct{ value int8 }

func (inst *BIPUSH) FetchOperands(reader *instruction.ByteCodeReader) {
	inst.value = reader.ReadInt8()
}

func (inst *BIPUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(inst.value))
}

// short sign-extended to int
type SIPUSH struct{ value int16 }

func (inst *SIPUSH) FetchOperands(reader *instruction.ByteCodeReader) {
	inst.value = reader.ReadInt16()
}

func (inst *SIPUSH) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(inst.value))
}
