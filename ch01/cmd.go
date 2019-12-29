package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	class       string
	args        []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage : %s [-optioins] class [args...]\n", os.Args[0])
}

func startJVM(cmd *Cmd) {
	fmt.Printf("start JVM, class : %s, classpath : %s, args : %v \n", cmd.class, cmd.cpOption, cmd.args)
}

func main() {
	cmd := parseCmd()

	if cmd.helpFlag {
		printUsage()
	}
	if cmd.versionFlag {
		fmt.Println("Version 1.0.0")
	}
	if cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}
