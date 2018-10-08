package main

import (
	"fmt"
	"jvmgo/classpath"
	"jvmgo/classfile"
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
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)

	className := strings.Replace(cmd.class, ".", "/", -1)
	data, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Error on load class %s %s\n", cmd.class, err)
	} else {
		fmt.Printf("class data: %v\n", data)
	}

	cf, err := classfile.ParseClassBytes(data)
	fmt.Println(cf)
}
