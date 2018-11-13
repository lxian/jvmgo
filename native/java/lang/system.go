package lang

import (
	"jvmgo/rtda"
	"jvmgo/native"
	"jvmgo/rtda/heap"
)

func init() {
	native.RegisterNativeMethod("java/lang/System", "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
}

func arraycopy(frame *rtda.Frame) {
	src := frame.LocalVars().GetRef(0)
	srcPos := frame.LocalVars().GetInt(1)
	dest := frame.LocalVars().GetRef(2)
	destPos := frame.LocalVars().GetInt(3)
	length := frame.LocalVars().GetInt(4)

	// nil checking
	if src == nil || dest == nil {
		panic("java.lang.NullPointerException")
	}

	// clz checking
	srcClz := src.Class()
	destClz := dest.Class()
	if !(srcClz.IsArray() && destClz.IsArray()) {
		panic("java.lang.ArrayStoreException")
	}
	if (srcClz.ComponentClass().IsPrimitive() || destClz.ComponentClass().IsPrimitive()) && srcClz != destClz {
		panic("java.lang.ArrayStoreException")
	}

	// index bounds checking
	if (src.ArrayLength() < srcPos + length) || (dest.ArrayLength() < destPos + length) {
		panic("java.lang.IndexOutOfBoundsException")
	}

	heap.ArrayCopy(src, dest, srcPos, destPos, length)
}
