package lang

import (
	"jvmgo/native"
	"jvmgo/rtda"
	"math"
)

func init() {
	native.RegisterNativeMethod("java/lang/Double", "doubleToRawLongBits", "(D)J", doubleToRawLongBits)
	native.RegisterNativeMethod("java/lang/Double", "longBitsToDouble", "(J)D", longBitsToDouble)
}

func doubleToRawLongBits(frame *rtda.Frame) {
	d := frame.LocalVars().GetDouble(0)
	frame.OperandStack().PushLong(int64(math.Float64bits(d)))
}

func longBitsToDouble(frame *rtda.Frame) {
	l := frame.LocalVars().GetLong(0)
	frame.OperandStack().PushDouble(math.Float64frombits(uint64(l)))
}
