package main

import (
	"flag"
	"fmt"
	"image/png"
	"io"
	"os"
)

func main() {
	// TODO Add a CLI with a help message and options. Some options can include:
	//
	//		- Changing colors (e.g. `-red "#FF8888")
	//		- Setting output path
	flag.Parse()
	args := flag.Args()

	if len(args) != 2 {
		flag.Usage()
		os.Exit(1)
	}

	// TODO Use an `io.Reader` in `lex` so that conversion to string isn't necessary.
	src, err := readInput(args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	art := parse(lex(src))

	m := art.asImage()

	// TODO Allow setting output path.
	dst, err := os.Create(args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	defer dst.Close()

	if err := png.Encode(dst, m); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
}

func readInput(input string) (string, error) {
	const megabyte = 1 << 20
	var r io.Reader
	buf := make([]byte, megabyte)
	if input == "-" {
		r = os.Stdin
	} else {
		f, err := os.Open(input)
		if err != nil {

			return "", err
		}
		defer f.Close()
		r = f
	}
	n, err := r.Read(buf)
	if err != nil && err != io.EOF {
		return "", err
	}
	return string(buf[:n]), nil
}
