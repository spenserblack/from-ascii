package main

import (
	"image/color"
	"iter"
	"strconv"
	"strings"
)

// lex creates a node iterator.
//
// It should never fail to return a node until the text is exhausted.
func lex(text string) iter.Seq[node] {
	return func(yield func(node) bool) {
		for text != "" {
			if strings.HasPrefix(text, "\n") {
				if !yield(newline(unix)) {
					return
				}
				text = text[1:]
			} else if strings.HasPrefix(text, "\r\n") {
				if !yield(newline(windows)) {
					return
				}
				text = text[2:]
			} else if strings.HasPrefix(text, control[0]) {
				n, rem := lexControl(text)
				if n == nil {
					// NOTE Failed to parse control, so falling back to treating as an invisible
					// character.
					n = invisible(text[0:1])
				}
				if !yield(n) {
					return
				}
				text = rem
			} else  {
				if !yield(character{Value: text[0]}) {
					return
				}
				text = text[1:]
			}
		}
	}
}

// lexControl lexes one of the controls for styling.
//
// It returns a node and the string remainder. If it fails to parse, then the node is
// nil and the remainder is the original text.
func lexControl(text string) (n node, remainder string) {
	style := text[len(control[0]):]
	end := strings.Index(style, control[1])
	if end == -1 {
		return nil, text
	}
	style = style[:end]
	remainder = text[len(style)+len(control[0])+len(control[1]):]
	return lexStyle(style), remainder
}

// lexStyle lexes a style string.
//
// It returns a style node.
func lexStyle(style string) styleNode {
	var n styleNode
	n.literal = control[0] + style + control[1]

	instructions := strings.Split(style, ";")

	// NOTE Everyone hates labels and say that it makes code unreadable, but I don't care.
	//		Refactor if you want.
ReadLoop:
	for i := 0; i < len(instructions); i++ {
		instruction := instructions[i]
		if instruction == reset {
			n.reset = true
			continue ReadLoop
		}
		for fg, c := range fgColorMap {
			if fg == instruction {
				n.fg = c
				n.reset = false
				continue ReadLoop
			}
		}
		for bg, c := range bgColorMap {
			if bg == instruction {
				n.bg = c
				n.reset = false
				continue ReadLoop
			}
		}

		// NOTE Either 256-color or true-color in use
		isAdvancedFg := instruction == fgAdvanced
		isAdvancedBg := instruction == bgAdvanced
		if isAdvancedFg || isAdvancedBg {
			flag := getAt(instructions, i+1)
			if flag == flag256 {
				// NOTE We're reading the next 2 instructions, so we'll skip on the next iteration.
				i += 2
				lookup := getAt(instructions, i)
				if c, ok := colorMap256[lookup]; ok {
					if isAdvancedFg {
						n.fg = c
						n.reset = false
					} else {
						n.bg = c
						n.reset = false
					}
				}
			} else if flag == flagTrueColor {
				n.reset = false
				i += 2
				r := getAtDefault(instructions, i, "0")
				g := getAtDefault(instructions, i+1, "0")
				b := getAtDefault(instructions, i+2, "0")
				// NOTE Skip the (at most) 3 color channels
				i += 3
				c := color.RGBA{A: 0xFF}
				// NOTE If any of these don't parse to [0, 256), that's undefined behavior.
				pairs := [3]struct {
					channel *uint8
					raw     string
				}{
					{
						channel: &c.R,
						raw:     r,
					},
					{
						channel: &c.G,
						raw:     g,
					},
					{
						channel: &c.B,
						raw:     b,
					},
				}
				for _, pair := range pairs {
					if n, err := strconv.Atoi(pair.raw); err == nil && n >= 0 && n <= 0xFF {
						*pair.channel = uint8(n)
					}
				}

				if isAdvancedFg {
					n.fg = c
				} else {
					n.bg = c
				}
			}
		}
	}

	return n
}

// getAt gets the value at a slice's index, or the zero value.
func getAt[T any](slice []T, index int) (item T) {
	if index < len(slice) {
		return slice[index]
	}
	return
}

// getAtDefault is like getAt, but returns a default if the index is out of bounds.
func getAtDefault[T any](slice []T, index int, d T) T {
	if index < len(slice) {
		return slice[index]
	}
	return d
}
