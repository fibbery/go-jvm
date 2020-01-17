package main

import (
	"fmt"
	"github.com/fibbery/go-jvm/classpath"
	"os"
	"strings"
)

func main() {
	cmd := parseCmd()

	if cmd.helpFlag {
		printUsage()
		os.Exit(1)
	}
	if cmd.versionFlag {
		fmt.Println("Version 1.0.0")
		os.Exit(1)
	}
	if cmd.mainClass == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Println("Start JVM .... ")
	cp := classpath.Parse(cmd.jreOption, cmd.cpOption)
	fmt.Printf("classpath:%v mainClass:%v args:%v\n", cp, cmd.mainClass, cmd.args)
	bytes, _, e := cp.ReadClass(strings.ReplaceAll(cmd.mainClass, ".", "/"))
	if e != nil {
		fmt.Printf("can't load main mainClass : %v\n", cmd.mainClass)
		return
	}
	fmt.Printf("class data : %v \n", bytes)
}
