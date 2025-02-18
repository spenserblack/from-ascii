package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
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

	var encoder func(io.Writer, image.Image) error
	imgFormat := formatFlag.String()
	if imgFormat == "auto" {
		imgFormat = strings.TrimPrefix(filepath.Ext(args[1]), ".")
	}

	switch imgFormat {
	case "png":
		encoder = png.Encode
	case "jpeg", "jpg":
		encoder = encodeJpeg
	case "gif":
		encoder = encodeGif
	default:
		fmt.Fprintf(os.Stderr, "unsupported format: %s\n", imgFormat)
		return
	}

	if err := encoder(dst, m); err != nil {
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

// encodeJpeg encodes a JPEG with the default options. This allows for the signature to
// match the other encoding functions.
func encodeJpeg(w io.Writer, m image.Image) error {
	return jpeg.Encode(w, m, nil)
}

// encodeGif encodes a GIF with the default options. This allows for the signature to
// match the other encoding functions.
func encodeGif(w io.Writer, m image.Image) error {
	return gif.Encode(w, m, nil)
}
