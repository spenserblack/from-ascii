package main

import (
	"errors"
	"slices"
	"strings"
)

type format struct {
	s string
}

// String implements the fmt.Stringer interface.
func (f format) String() string {
	return f.s
}

var validFormats = []string{"auto", "png", "jpeg", "jpg", "gif"}

// errInvalidFormat is returned when an invalid format is passed.
var errInvalidFormat = errors.New("must be one of: " + strings.Join(validFormats, ", "))

// Set implements the flag.Value interface.
func (f *format) Set(text string) error {
	if !slices.Contains(validFormats, text) {
		return errInvalidFormat
	}
	f.s = text
	return nil
}
