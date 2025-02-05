package main

import "image/color"

// node is an item in the ASCII art.
type node interface {
	// Literal returns the literal representation of the node.
	Literal() string
}

// character is a single character in the ASCII art.
type character struct {
	// Value is the character.
	Value byte
}

// Literal returns the literal representation of the character.
func (c character) Literal() string {
	return string(c.Value)
}

// String returns the character as a string.
func (c character) String() string {
	return string(c.Value)
}

// invisible is used to handle unknown characters that shouldn't be displayed.
type invisible string

// Literal returns the string that was considered invisible.
func (i invisible) Literal() string {
	return string(i)
}

// String returns an empty string (because the character is invisible).
func (i invisible) String() string {
	return ""
}

// newline is a newline character.
type newline bool

const (
	unix    newline = false
	windows newline = true
)

// Literal returns the literal representation of the newline.
func (n newline) Literal() string {
	if n {
		return "\r\n"
	}
	return "\n"
}

// String returns the newline value.
func (n newline) String() string {
	if n {
		return `\r\n`
	}
	return `\n`
}

// styleNode marks a change in styling.
type styleNode struct {
	// fg is the foreground color.
	fg color.Color
	// bg is the background color.
	bg color.Color
	// reset means that the style has been reset to whatever the defaults are. If this is
	// true, then the foreground and background colors should be ignored.
	reset bool
	// literal is the literal string that generated the node.
	literal string
}

// Literal returns the literal representation of the styleNode.
func (n styleNode) Literal() string {
	return n.literal
}
