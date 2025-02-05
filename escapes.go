package main

import (
	"image/color"
	"strconv"
)

// NOTE: See https://en.wikipedia.org/wiki/ANSI_escape_code#Colors

const (
	reset     = "0"
	fgBlack   = "30"
	bgBlack   = "40"
	fgRed     = "31"
	bgRed     = "41"
	fgGreen   = "32"
	bgGreen   = "42"
	fgYellow  = "33"
	bgYellow  = "43"
	fgBlue    = "34"
	bgBlue    = "44"
	fgMagenta = "35"
	bgMagenta = "45"
	fgCyan    = "36"
	bgCyan    = "46"
	fgWhite   = "37"
	bgWhite   = "47"

	fgBrightBlack   = "90"
	bgBrightBlack   = "100"
	fgBrightRed     = "91"
	bgBrightRed     = "101"
	fgBrightGreen   = "92"
	bgBrightGreen   = "102"
	fgBrightYellow  = "93"
	bgBrightYellow  = "103"
	fgBrightBlue    = "94"
	bgBrightBlue    = "104"
	fgBrightMagenta = "95"
	bgBrightMagenta = "105"
	fgBrightCyan    = "96"
	bgBrightCyan    = "106"
	fgBrightWhite   = "97"
	bgBrightWhite   = "107"

	flag256       = "5"
	flagTrueColor = "2"
	fgAdvanced    = "38"
	bgAdvanced    = "48"
)

var (
	control = [2]string{"\033[", "m"}
)

// fgColorMap maps simple foreground color codes to a Go struct.
var fgColorMap = map[string]color.Color{
	fgBlack:         color.Black,
	fgRed:           color.RGBA{R: 0xAA, G: 0, B: 0, A: 0xFF},
	fgGreen:         color.RGBA{R: 0, G: 0xAA, B: 0, A: 0xFF},
	fgYellow:        color.RGBA{R: 0xAA, G: 0xAA, B: 0, A: 0xFF},
	fgBlue:          color.RGBA{R: 0, G: 0, B: 0xAA, A: 0xFF},
	fgMagenta:       color.RGBA{R: 0xAA, G: 0, B: 0xAA, A: 0xFF},
	fgCyan:          color.RGBA{R: 0, G: 0xAA, B: 0xAA, A: 0xFF},
	fgWhite:         color.RGBA{R: 0xAA, G: 0xAA, B: 0xAA, A: 0xFF},
	fgBrightBlack:   color.Gray{0x55},
	fgBrightRed:     color.RGBA{R: 0xFF, G: 0x55, B: 0x55, A: 0xFF},
	fgBrightGreen:   color.RGBA{R: 0x55, G: 0xFF, B: 0x55, A: 0xFF},
	fgBrightYellow:  color.RGBA{R: 0xFF, G: 0xFF, B: 0x55, A: 0xFF},
	fgBrightBlue:    color.RGBA{R: 0x55, G: 0x55, B: 0xFF, A: 0xFF},
	fgBrightMagenta: color.RGBA{R: 0xFF, G: 0x55, B: 0xFF, A: 0xFF},
	fgBrightCyan:    color.RGBA{R: 0x55, G: 0xFF, B: 0xFF, A: 0xFF},
	fgBrightWhite:   color.White,
}

// bgColorMap maps simple background color codes to a Go struct.
var bgColorMap = map[string]color.Color{
	bgBlack:         color.Black,
	bgRed:           color.RGBA{R: 0xAA, G: 0, B: 0, A: 0xFF},
	bgGreen:         color.RGBA{R: 0, G: 0xAA, B: 0, A: 0xFF},
	bgYellow:        color.RGBA{R: 0xAA, G: 0xAA, B: 0, A: 0xFF},
	bgBlue:          color.RGBA{R: 0, G: 0, B: 0xAA, A: 0xFF},
	bgMagenta:       color.RGBA{R: 0xAA, G: 0, B: 0xAA, A: 0xFF},
	bgCyan:          color.RGBA{R: 0, G: 0xAA, B: 0xAA, A: 0xFF},
	bgWhite:         color.RGBA{R: 0xAA, G: 0xAA, B: 0xAA, A: 0xFF},
	bgBrightBlack:   color.Gray{0x55},
	bgBrightRed:     color.RGBA{R: 0xFF, G: 0x55, B: 0x55, A: 0xFF},
	bgBrightGreen:   color.RGBA{R: 0x55, G: 0xFF, B: 0x55, A: 0xFF},
	bgBrightYellow:  color.RGBA{R: 0xFF, G: 0xFF, B: 0x55, A: 0xFF},
	bgBrightBlue:    color.RGBA{R: 0x55, G: 0x55, B: 0xFF, A: 0xFF},
	bgBrightMagenta: color.RGBA{R: 0xFF, G: 0x55, B: 0xFF, A: 0xFF},
	bgBrightCyan:    color.RGBA{R: 0x55, G: 0xFF, B: 0xFF, A: 0xFF},
	bgBrightWhite:   color.White,
}

// colorMap256 is the 256-color lookup table.
var colorMap256 = map[string]color.Color{
	"0":  color.Black,
	"1":  color.RGBA{R: 0xAA, G: 0, B: 0, A: 0xFF},
	"2":  color.RGBA{R: 0, G: 0xAA, B: 0, A: 0xFF},
	"3":  color.RGBA{R: 0xAA, G: 0xAA, B: 0, A: 0xFF},
	"4":  color.RGBA{R: 0, G: 0, B: 0xAA, A: 0xFF},
	"5":  color.RGBA{R: 0xAA, G: 0, B: 0xAA, A: 0xFF},
	"6":  color.RGBA{R: 0, G: 0xAA, B: 0xAA, A: 0xFF},
	"7":  color.RGBA{R: 0xAA, G: 0xAA, B: 0xAA, A: 0xFF},
	"8":  color.Gray{0x55},
	"9":  color.RGBA{R: 0xFF, G: 0x55, B: 0x55, A: 0xFF},
	"10": color.RGBA{R: 0x55, G: 0xFF, B: 0x55, A: 0xFF},
	"11": color.RGBA{R: 0xFF, G: 0xFF, B: 0x55, A: 0xFF},
	"12": color.RGBA{R: 0x55, G: 0x55, B: 0xFF, A: 0xFF},
	"13": color.RGBA{R: 0xFF, G: 0x55, B: 0xFF, A: 0xFF},
	"14": color.RGBA{R: 0x55, G: 0xFF, B: 0xFF, A: 0xFF},
	"15": color.White,
}

func init() {
	// NOTE Initialize colors 16-255 for the 256-color lookup table.
	// 6x6x6 cube of colors, plus dark-to-light in 24 steps.
	// NOTE 0x33 * 5 = 0xFF
	for r := 0; r <= 0xFF; r += 0x33 {
		for g := 0; g <= 0xFF; g += 0x33 {
			for b := 0; b <= 0xFF; b += 0x33 {
				c := color.RGBA{
					R: uint8(r),
					G: uint8(g),
					B: uint8(b),
					A: 0xFF,
				}
				// NOTE Max value = 231 (0xE7)
				key := strconv.Itoa(0x10 + ((r / 0x33) * 0x24) + ((g / 0x33) * 0x06) + (b / 0x33))
				colorMap256[key] = c
			}
		}
	}
	for step := 0; step <= 0x17; step++ {
		c := color.Gray{uint8((step * 0xFF) / 0x17)}
		key := strconv.Itoa(0xE8 + step)
		colorMap256[key] = c
	}
}
