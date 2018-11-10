package native

import "jvmgo/rtda"

type NativeMethod func(frame *rtda.Frame)

func EmptyNativeMethod(frame *rtda.Frame) {
}
