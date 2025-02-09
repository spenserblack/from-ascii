package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [flags] <INPUT> <OUTPUT>\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Positional arguments:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\tINPUT\tPath to the input file. Use '-' for stdin.\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\tOUTPUT\tPath to the output file.\n")
		flag.PrintDefaults()
	}
}
