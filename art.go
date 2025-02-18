package main

import (
	"image"
	"image/color"
	"image/draw"
)

// art is parsed ASCII art.
type art struct {
	// lines is an array of lines (rows) in the ASCII art.
	lines [][]node
	// width is the detected width of the art.
	width int
}

// height gets the number of lines (rows) in the ASCII art.
func (a art) height() int {
	return len(a.lines)
}

// asImage converts the art to an image.
func (a art) asImage() image.Image {
	var (
		bgColor color.Color = defaultBgColor
		fgColor color.Color = defaultFgColor
	)
	rect := image.Rect(0, 0, a.width, a.height())
	background := image.NewRGBA(rect)
	foreground := image.NewRGBA(rect)

	for y, line := range a.lines {
		x := 0
		for _, node := range line {
			switch v := node.(type) {
			case character:
				alpha, ok := characterAlphas[v.Value]
				if !ok {
					alpha = 0xFF
				}
				background.Set(x, y, bgColor)
				foreground.Set(x, y, setAlpha(fgColor, alpha))
				x += 1
			case styleNode:
				if v.reset {
					bgColor = defaultBgColor
					fgColor = defaultFgColor
				}
				// NOTE It shouldn't be possible for reset to be true and fg or bg to be
				//		not nil.
				if v.bg != nil {
					bgColor = v.bg
				}
				if v.fg != nil {
					fgColor = v.fg
				}
			default:
				panic("unreachable")
			}
		}
	}

	draw.Draw(background, rect, foreground, image.Point{X: 0, Y: 0}, draw.Over)

	return background
}

// setAlpha changes a color's opacity
func setAlpha(c color.Color, alpha uint8) color.Color {
	// NOTE Colors are in the range [0, 0xFFFF].
	r, g, b, _ := c.RGBA()
	return color.NRGBA{
		R: uint8(r >> 8),
		G: uint8(g >> 8),
		B: uint8(b >> 8),
		A: alpha,
	}
}
