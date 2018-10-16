package stack

import (
	"github.com/derekparker/delve/pkg/dwarf/frame"
	"jvmgo/instruction"
	"jvmgo/rtda"
)

// POP
type POP struct {
	instruction.NoOperandsInstruction
}
type POP2 struct {
	instruction.NoOperandsInstruction
}

func (*POP) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
}

func (*POP2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
	frame.OperandStack().PopSlot()
}

// DUP
type DUP struct {
	instruction.NoOperandsInstruction
}
type DUP_X1 struct {
	instruction.NoOperandsInstruction
}
type DUP_X2 struct {
	instruction.NoOperandsInstruction
}

func (*DUP) Execute(frame *rtda.Frame) {
	slot := frame.OperandStack().PeekSlot()
	frame.OperandStack().PushSlot(slot)
}

func (*DUP_X1) Execute(frame *rtda.Frame) {
	slotToDup := frame.OperandStack().PopSlot()
	slot1 := frame.OperandStack().PopSlot()
	frame.OperandStack().PushSlot(slotToDup)
	frame.OperandStack().PushSlot(slot1)
	frame.OperandStack().PushSlot(slotToDup)
}

func (*DUP_X2) Execute(frame *rtda.Frame) {
	slotToDup := frame.OperandStack().PopSlot()
	slot1 := frame.OperandStack().PopSlot()
	slot2 := frame.OperandStack().PopSlot()
	frame.OperandStack().PushSlot(slotToDup)
	frame.OperandStack().PushSlot(slot2)
	frame.OperandStack().PushSlot(slot1)
	frame.OperandStack().PushSlot(slotToDup)
}

type DUP2 struct {
	instruction.NoOperandsInstruction
}
type DUP2_X1 struct {
	instruction.NoOperandsInstruction
}
type DUP2_X2 struct {
	instruction.NoOperandsInstruction
}

func (*DUP2) Execute(frame *rtda.Frame) {
	slot_h := frame.OperandStack().PopSlot()
	slot_l := frame.OperandStack().PopSlot()
	frame.OperandStack().PushSlot(slot_l)
	frame.OperandStack().PushSlot(slot_h)
	frame.OperandStack().PushSlot(slot_l)
	frame.OperandStack().PushSlot(slot_h)
}

func (*DUP2_X1) Execute(frame *rtda.Frame) {
	slot_h := frame.OperandStack().PopSlot()
	slot_l := frame.OperandStack().PopSlot()
	slot_1 := frame.OperandStack().PopSlot()
	frame.OperandStack().PushSlot(slot_l)
	frame.OperandStack().PushSlot(slot_h)
	frame.OperandStack().PushSlot(slot_1)
	frame.OperandStack().PushSlot(slot_l)
	frame.OperandStack().PushSlot(slot_h)
}

func (*DUP2_X2) Execute(frame *rtda.Frame) {
	slot_h := frame.OperandStack().PopSlot()
	slot_l := frame.OperandStack().PopSlot()
	slot_1 := frame.OperandStack().PopSlot()
	slot_2 := frame.OperandStack().PopSlot()
	frame.OperandStack().PushSlot(slot_l)
	frame.OperandStack().PushSlot(slot_h)
	frame.OperandStack().PushSlot(slot_2)
	frame.OperandStack().PushSlot(slot_1)
	frame.OperandStack().PushSlot(slot_l)
	frame.OperandStack().PushSlot(slot_h)
}

// Swap
type SWAP struct {
	instruction.NoOperandsInstruction
}

func (*SWAP) Execute(frame *rtda.Frame) {
	slot_1 := frame.OperandStack().PopSlot()
	slot_2 := frame.OperandStack().PopSlot()
	frame.OperandStack().PushSlot(slot_1)
	frame.OperandStack().PushSlot(slot_2)
}
