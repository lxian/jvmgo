package main

import (
	"jvmgo/rtda"
	"jvmgo/rtda/heap"
	"jvmgo/classpath"
	"strings"
)

type JVM struct {
	cmd *Cmd
	classLoader *heap.ClassLoader
	mainThread *rtda.Thread
}

func newJVM(cmd *Cmd) *JVM {
	//fmt.Printf("classpath:%s class: %s args: %v\n", cmd.cpOption, cmd.class, cmd.args)
	classPath := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLaoder(classPath, cmd.verboseClassFlag)
	mainThread := rtda.NewThread()

	return &JVM{
		cmd: cmd,
		classLoader: classLoader,
		mainThread: mainThread,
	}
	//fmt.Println("Found main method", mainMethod.Name(), mainMethod.Descriptor())
}

func (jvm *JVM) start() {
	jvm.initVM()
	jvm.execMain()
}

func (jvm *JVM) initVM() {
	vmClz := jvm.classLoader.LoadClass("sun/misc/VM");
	vmInitMethod := vmClz.FindInitMethod()

	thread := jvm.mainThread
	frame := rtda.NewFrame(thread, vmInitMethod)
	thread.PushFrame(frame)

	interpret(thread, jvm.cmd.verboseInstFlag)
}

func (jvm *JVM) execMain() {
	mainClassName := strings.Replace(jvm.cmd.class, ".", "/", -1)
	mainClass := jvm.classLoader.LoadClass(mainClassName)
	mainMethod := findMainMethod(mainClass)

	thread := jvm.mainThread
	frame := rtda.NewFrame(thread, mainMethod)
	thread.PushFrame(frame)

	interpret(thread, jvm.cmd.verboseInstFlag)
}

func findMainMethod(class *heap.Class) *heap.Method {
	for _, method := range class.Methods() {
		if method.Name() == "main" && method.Descriptor() == "([Ljava/lang/String;)V" {
			return method
		}
	}
	return nil
}
