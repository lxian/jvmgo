package main

import (
	"fmt"
	"jvmgo/classpath"
	"jvmgo/rtda/heap"
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

	classLoader := heap.NewClassLaoder(classPath, cmd.verboseClassFlag)
	class := classLoader.LoadClass(className)
	mainMethod := findMainMethod(class)
	fmt.Println("Found main method", mainMethod.Name(), mainMethod.Descriptor())

	interpret(mainMethod, cmd.verboseInstFlag)
}

func findMainMethod(class *heap.Class) *heap.Method {
	for _, method := range class.Methods() {
		if method.Name() == "main" && method.Descriptor() == "([Ljava/lang/String;)V" {
			return method
		}
	}
	return nil
}
