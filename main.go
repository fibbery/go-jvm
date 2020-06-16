package main

import (
	"fmt"
	"github.com/fibbery/go-jvm/classfile"
	"os"
	"strings"

	"github.com/fibbery/go-jvm/classpath"
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
	cf := loadClass(strings.ReplaceAll(cmd.mainClass, ".", "/"), cp)
	if cf != nil {
		printClassInfo(cf)
	}
}

func loadClass(className string, cp *classpath.Classpath) *classfile.ClassFile {
	bytes, _, err := cp.ReadClass(className)
	if err != nil {
		panic(err)
	}
	cf, err := classfile.Parse(bytes)
	if err != nil {
		panic(err)
	}
	return cf
}

func printClassInfo(cf *classfile.ClassFile) {
	fmt.Printf("version : %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("constant count : %v\n", len(cf.ConstPool()))
	fmt.Printf("access flags : 0x%x\n", cf.AccessFlag())
	fmt.Printf("this class : %v\n", strings.ReplaceAll(cf.ThisClass(),"/", "."))
	fmt.Printf("super class : %v\n", strings.ReplaceAll(cf.SuperClass(),"/", "."))
	fmt.Printf("fields count : %v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("  %s\n", f.Name())
	}
	fmt.Printf("methods count: %v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("  %s\n", m.Name())
	}
}
