package reference

import (
	"jvmgo/instruction"
	"jvmgo/native"
	_ "jvmgo/native/java/lang"
	_ "jvmgo/native/sun/misc"
	"jvmgo/rtda"
)

type INVOKE_NATIVE struct {
	instruction.NoOperandsInstruction
}

func (inst *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	clz := frame.Method().Class()
	curMethod := frame.Method()

	nativeMethod := native.FindNativeMethod(clz.Name(), curMethod.Name(), curMethod.Descriptor())
	nativeMethod(frame)
}
