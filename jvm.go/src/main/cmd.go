package main

import "flag"
import "fmt"
import "os"

type Cmd struct {
	helpFlag    bool
	versionFlag bool
	cpOption    string
	xjreOption  string
	class       string
	args        []string
}

func parseCmd() *Cmd {

	cmd := &Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "? ", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print verion and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.xjreOption, "XjreOption", "", "path to jre")

	flag.Parse()

	args := flag.Args()

	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd

}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.Xjreoption, cmd, cpOption)
	fme.Printf()
	fmt.Printf("classpath:%s class:%s args:%v\n", cp, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)

	if err != nil {
		fmt.Printf("Could not find or load main class %\n", cmd.class)
		return
	}

	fmt.Printf("class data :%v\n", classData)
}
