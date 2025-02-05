package main

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		name   string
		nodes  []node
		width  int
		height int
	}{
		{
			name: "It sets proper dimensions",
			nodes: []node{
				character{Value: '#'}, unix,
				character{Value: '#'}, character{Value: '#'}, unix,
				character{Value: '#'}, character{Value: '#'}, character{Value: '#'}, unix,
				character{Value: '#'}, character{Value: '#'}, unix,
				character{Value: '#'},
			},
			width:  3,
			height: 5,
		},
		{
			name: "It sets ignores final newlines",
			nodes: []node{
				character{Value: '#'}, unix,
				character{Value: '#'}, character{Value: '#'}, unix,
				character{Value: '#'}, character{Value: '#'}, character{Value: '#'}, unix,
				character{Value: '#'}, character{Value: '#'}, unix,
				character{Value: '#'}, unix,
			},
			width:  3,
			height: 5,
		},
		{
			name:   "It handles empty inputs",
			nodes:  []node{},
			width:  0,
			height: 0,
		},
		{
			name:   "It ignores invisible characters",
			nodes:  []node{character{Value: '#'}, invisible("\033"), character{Value: '#'}},
			width:  2,
			height: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := parse(tt.nodes)
			if a.width != tt.width {
				t.Errorf("a.width = %d, want %d", a.width, tt.width)
			}
			if got := a.height(); got != tt.height {
				t.Errorf("a.height() = %d, want %d", got, tt.height)
			}
		})
	}
}
