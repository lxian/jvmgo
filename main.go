package main

import (
	"fmt"
	"jvmgo/classfile"
	"jvmgo/classpath"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class: %s args: %v\n", cmd.cpOption, cmd.class, cmd.args)
	classPath := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classFile := loadClass(className, classPath)

	mainMethod := findMainMethod(classFile)
	fmt.Println("Found main method", mainMethod.Name(), mainMethod.Descriptor())

	interpret(mainMethod)
	//frame := rtda.NewFrame(100, 100, thr)
	//frame.LocalVars().SetDouble(1, 1.234)
	//fmt.Println(frame.LocalVars().GetDouble(1))
	//frame.OperandStack().PushDouble(1.234)
	//fmt.Println(frame.OperandStack().PopDouble())
}

func findMainMethod(classFile *classfile.ClassFile) *classfile.MemberInfo {
	for _, methodInfo := range classFile.Methods() {
		if methodInfo.Name() == "main" && methodInfo.Descriptor() == "([Ljava/lang/String;)V" {
			return methodInfo
		}
	}
	return nil
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	data, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Error on load class %s %s\n", className, err)
	} else {
		fmt.Printf("class data: %v\n", data)
	}

	cf, err := classfile.ParseClassBytes(data)
	if err != nil {
		panic(err)
	}
	return cf
}
