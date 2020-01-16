package main

import (
	"fmt"
	"os"
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
	if cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("start JVM, class : %s, classpath : %s, args : %v \n", cmd.class, cmd.cpOption, cmd.args)
}



