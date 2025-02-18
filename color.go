package main

import (
	"errors"
	"fmt"
	"image/color"
	"strconv"
)

// cliColor is a color that can be unmarshaled from a CLI flag in hex format.
type cliColor struct {
	color.NRGBA
}

var errInvalidColor = errors.New(`invalid color format; must be in the format "RRGGBB" or "RRGGBBAA"`)

// String implements the fmt.Stringer interface.
func (c cliColor) String() string {
	return fmt.Sprintf("%02X%02X%02X%02X", c.R, c.G, c.B, c.A)
}

// Set implements the flag.Value interface.
func (c *cliColor) Set(text string) error {
	l := len(text)
	if l != 6 && l != 8 {
		return errInvalidColor
	}

	c.A = 0xFF

	if r, err := strconv.ParseUint(string(text[0:2]), 16, 8); err != nil {
		return err
	} else {
		c.R = uint8(r)
	}
	if g, err := strconv.ParseUint(string(text[2:4]), 16, 8); err != nil {
		return err
	} else {
		c.G = uint8(g)
	}
	if b, err := strconv.ParseUint(string(text[4:6]), 16, 8); err != nil {
		return err
	} else {
		c.B = uint8(b)
	}
	if l == 8 {
		if a, err := strconv.ParseUint(string(text[6:8]), 16, 8); err != nil {
			return err
		} else {
			c.A = uint8(a)
		}
	}
	return nil
}
