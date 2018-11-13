package lang

import (
	"jvmgo/rtda"
	"math"
	"jvmgo/native"
)

func init() {
	native.RegisterNativeMethod("java/lang/Double", "doubleToRawIntBits", "(D)L", doubleToRawIntBits)
}

func doubleToRawIntBits(frame *rtda.Frame) {
	d := frame.LocalVars().GetDouble(0)
	frame.OperandStack().PushLong(int64(math.Float64bits(d)))
}
