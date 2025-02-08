package main

import (
	"image/png"
	"io"
	"log"
	"os"
)

func main() {
	// TODO Add a CLI with a help message and options. Some options can include:
	//
	//		- Changing colors (e.g. `-red "#FF8888")
	//		- Setting output path
	args := os.Args[1:]

	// TODO Use an `io.Reader` in `lex` so that conversion to string isn't necessary.
	src, err := readInput(args[0])
	if err != nil {
		log.Fatal(err)
	}
	art := parse(lex(src))

	m := art.asImage()

	// TODO Allow setting output path.
	dst, err := os.Create("ascii-art.png")
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	if err := png.Encode(dst, m); err != nil {
		log.Fatal(err)
	}
}

func readInput(input string) (string, error) {
	var r io.Reader
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
	b, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
