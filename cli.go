package main

import (
	"flag"
	"fmt"
	"image/color"
	"os"
	"strings"
)

var (
	// defaultBgColor is the default background when the background is reset or unspecified.
	defaultBgColor = cliColor{color.NRGBA{R: 0x00, G: 0x00, B: 0x00, A: 0xFF}}
	// defaultFgColor is the default foreground when the background is reset or unspecified.
	defaultFgColor = cliColor{color.NRGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}}
	// formatFlag is the format to output the image in.
	formatFlag = format{"auto"}
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [flags] <INPUT> <OUTPUT>\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Positional arguments:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\tINPUT\tPath to the input file. Use '-' for stdin.\n")
		fmt.Fprintf(flag.CommandLine.Output(), "\tOUTPUT\tPath to the output file.\n")
		flag.PrintDefaults()
	}
	flag.Var(&defaultBgColor, "bg", "Background color in hex format (RRGGBB or RRGGBBAA).")
	flag.Var(&defaultFgColor, "fg", "Foreground color in hex format (RRGGBB or RRGGBBAA).")
	flag.Var(&formatFlag, "format", "Output format ["+strings.Join(validFormats, "|")+"].")
}
