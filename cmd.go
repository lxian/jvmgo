package main

import "flag"
import "fmt"
import "os"

type Cmd struct {
	helpFlag         bool
	versionFlag      bool
	XjreOption       string
	cpOption         string
	class            string
	args             []string
	verboseClassFlag bool
	verboseInstFlag  bool
}

func parseCmd() *Cmd {
	cmd := new(Cmd)

	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "JRE base dir")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose:class", false, "print class info on class load")
	flag.BoolVar(&cmd.verboseInstFlag, "verbose:inst", false, "print instruction on execution")
	flag.Parse()

	if args := flag.Args(); len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func printUsage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}
