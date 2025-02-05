package main

// characterAlphas is a map of alpha values for each character. An unspecified character
// should be 100% opaque.
var characterAlphas = map[byte]uint8{
	' ':  0,
	'\'': 0x40,
	'"':  0x40,
	':':  0x80,
	';':  0x80,
}
