package extended

import (
	"fmt"
	"jvmgo/instruction"
	"jvmgo/rtda"
)

type WIDE struct {
	opcode     uint8
	index      uint16
	constValue int16
}

func (wide *WIDE) FetchOperands(reader *instruction.ByteCodeReader) {
	wide.opcode = reader.ReadUint8()
	wide.index = reader.ReadUint16()
	if wide.opcode == 0x84 {
		wide.constValue = reader.ReadInt16()
	}
}

func (wide *WIDE) Execute(frame *rtda.Frame) {
	index := uint(wide.index)
	switch wide.opcode {
	case 0x15:
		frame.OperandStack().PushInt(frame.LocalVars().GetInt(index))
	case 0x16:
		frame.OperandStack().PushLong(frame.LocalVars().GetLong(index))
	case 0x17:
		frame.OperandStack().PushFloat(frame.LocalVars().GetFloat(index))
	case 0x18:
		frame.OperandStack().PushDouble(frame.LocalVars().GetDouble(index))
	case 0x19:
		frame.OperandStack().PushRef(frame.LocalVars().GetRef(index))
	case 0x36:
		frame.LocalVars().SetInt(index, frame.OperandStack().PopInt())
	case 0x37:
		frame.LocalVars().SetLong(index, frame.OperandStack().PopLong())
	case 0x38:
		frame.LocalVars().SetFloat(index, frame.OperandStack().PopFloat())
	case 0x39:
		frame.LocalVars().SetDouble(index, frame.OperandStack().PopDouble())
	case 0x3a:
		frame.LocalVars().SetRef(index, frame.OperandStack().PopRef())
	case 0x84:
		v := frame.LocalVars().GetInt(index)
		v += int32(wide.constValue)
		frame.LocalVars().SetInt(index, v)
	case 0xa9: // ret
		fmt.Print("wide ret")
	}
}
