package lang

import (
	"jvmgo/native"
	"jvmgo/rtda"
	"fmt"
)

func init() {
	native.RegisterNativeMethod("java/lang/Throwable", "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

type StackTraceElement struct {
	fileName string
	className string
	methodName string
	lineNumber int
}

func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
	method := frame.Method()
	clz := frame.Method().Class()
	return &StackTraceElement{
		fileName: clz.SourceFileName(),
		className: clz.Name(),
		methodName: method.Name(),
		lineNumber: method.GetLineNumber(frame.Thread().PC()),
	}
}

func (self *StackTraceElement) String() string {
	return fmt.Sprintf("%s.%s(%s:%d)",
		self.className, self.methodName, self.fileName, self.lineNumber)
}

func fillInStackTrace(frame *rtda.Frame) {
	this := frame.LocalVars().GetRef(0)

	frame.OperandStack().PushRef(this)
}

