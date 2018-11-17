package lang

import (
	"fmt"
	"jvmgo/native"
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
)

func init() {
	native.RegisterNativeMethod("java/lang/Throwable", "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

type StackTraceElement struct {
	fileName   string
	className  string
	methodName string
	lineNumber int
}

func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
	method := frame.Method()
	clz := frame.Method().Class()
	return &StackTraceElement{
		fileName:   clz.SourceFileName(),
		className:  clz.Name(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.NextPC() - 1),
	}
}

func (stackTraceEle *StackTraceElement) String() string {
	return fmt.Sprintf("%s.%s(%s:%d)",
		stackTraceEle.className, stackTraceEle.methodName, stackTraceEle.fileName, stackTraceEle.lineNumber)
}

func fillInStackTrace(frame *rtda.Frame) {
	this := frame.LocalVars().GetRef(0)

	stackIterator := rtda.NewStackIterator(frame.Thread())
	skipThrowableObjCreationFrames(this.Class(), stackIterator)
	traceElements := make([]*StackTraceElement, 0)
	for stackIterator.HasNext() {
		frame := stackIterator.Next()
		ele := createStackTraceElement(frame)
		traceElements = append(traceElements, ele)
	}
	this.SetExtra(traceElements)

	frame.OperandStack().PushRef(this)
}

func skipThrowableObjCreationFrames(expClz *heap.Class, iterator *rtda.StackIterator) {
	framesToSkip := superClzCount(expClz) + 1 + 1 // super class obj init + obj init + fillInStackTrace
	for iterator.HasNext() && framesToSkip > 0 {
		iterator.Next()
		framesToSkip -= 1
	}
}

func superClzCount(clz *heap.Class) int {
	cnt := 0
	for c := clz.SuperClass(); c != nil; c = c.SuperClass() {
		cnt += 1
	}
	return cnt
}
