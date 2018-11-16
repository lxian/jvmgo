package lang

import (
	"jvmgo/native"
	"jvmgo/rtda"
	"math"
)

func init() {
	native.RegisterNativeMethod("java/lang/Float", "floatToRawIntBits", "(F)I", floatToRawIntBits)
}

func floatToRawIntBits(frame *rtda.Frame) {
	f := frame.LocalVars().GetFloat(0)
	frame.OperandStack().PushInt(int32(math.Float32bits(f)))
}
