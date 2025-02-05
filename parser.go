package main

// parse creates a new art from lexed nodes.
func parse(nodes []node) art {
	a := art{
		lines: [][]node{{}},
		width: 0,
	}
	i := 0
	width := 0
	for _, n := range nodes {
		switch v := n.(type) {
		case invisible:
			continue
		case newline:
			a.lines = append(a.lines, []node{})
			i += 1
			if width > a.width {
				a.width = width
			}
			width = 0
		case styleNode:
			a.lines[i] = append(a.lines[i], v)
		case character:
			a.lines[i] = append(a.lines[i], v)
			width += 1
		default:
			panic("unreachable")
		}
	}
	// NOTE In case there weren't any newlines
	if width > a.width {
		a.width = width
	}
	// NOTE Ignore final newline
	height := len(a.lines)
	if height >= 1 {
		last := a.lines[height-1]
		if len(last) == 0 {
			a.lines = a.lines[:height-1]
		}
	}

	return a
}
