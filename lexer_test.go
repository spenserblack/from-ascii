package main

import (
	"image/color"
	"iter"
	"testing"
)

func TestLex(t *testing.T) {
	t.Run("lexes a character", func(t *testing.T) {
		nodes := lex("c")
		next, stop := iter.Pull(nodes)
		defer stop()
		node, ok := next()
		if !ok {
			t.Fatalf("Expected a node")
		}
		char, ok := node.(character)
		if !ok {
			t.Fatalf("Expected a character, got %T", node)
		}
		if v := char.Value; v != 'c' {
			t.Fatalf("Expected c, got %c", v)
		}
	})

	invTests := []struct {
		name string
		text string
		want invisible
	}{
		{
			name: "Lexes a control with no end as invisible",
			text: "\033[Foo",
			want: invisible("\033"),
		},
	}

	for _, tt := range invTests {
		t.Run(tt.name, func(t *testing.T) {
			nodes := lex(tt.text)
			next, stop := iter.Pull(nodes)
			defer stop()
			node, ok := next()
			if !ok {
				t.Fatalf("Expected a node")
			}
			inv, ok := node.(invisible)
			if !ok {
				t.Fatalf("Expected an invisible, got %T", node)
			}

			if inv != tt.want {
				t.Fatalf("Want %v, got %v", tt.want, inv)
			}
		})
	}

	newlineTests := []struct {
		name string
		text string
		want newline
	}{
		{
			name: "Lexes a unix newline",
			text: "\n",
			want: unix,
		},
		{
			name: "Lexes a windows newline",
			text: "\r\n",
			want: windows,
		},
	}

	for _, tt := range newlineTests {
		t.Run(tt.name, func(t *testing.T) {
			nodes := lex(tt.text)
			next, stop := iter.Pull(nodes)
			defer stop()
			node, ok := next()
			if !ok {
				t.Fatalf("Expected a node")
			}
			nl, ok := node.(newline)
			if !ok {
				t.Fatalf("Expected a newline, got %T", node)
			}

			if nl != tt.want {
				t.Fatalf("Want %v, got %v", tt.want, nl)
			}
		})
	}

	styleTests := []struct {
		name string
		text string
		want styleNode
	}{
		{
			name: "Lexes a black foreground color",
			text: "\033[30m",
			want: styleNode{
				fg: color.Black,
			},
		},
		{
			name: "Lexes a black background color",
			text: "\033[40m",
			want: styleNode{
				bg: color.Black,
			},
		},
		{
			name: "Lexes a reset",
			text: "\033[0m",
			want: styleNode{
				reset: true,
			},
		},
		{
			name: "Lexes 256-color black text from cube",
			text: "\033[38;5;16m",
			want: styleNode{
				fg: color.RGBA{0, 0, 0, 0xFF},
			},
		},
		{
			name: "Lexes 256-color red text from cube",
			text: "\033[38;5;196m",
			want: styleNode{
				fg: color.RGBA{0xFF, 0, 0, 0xFF},
			},
		},
		{
			name: "Lexes 256-color green text from cube",
			text: "\033[38;5;46m",
			want: styleNode{
				fg: color.RGBA{0, 0xFF, 0, 0xFF},
			},
		},
		{
			name: "Lexes 256-color blue text from cube",
			text: "\033[38;5;21m",
			want: styleNode{
				fg: color.RGBA{0, 0, 0xFF, 0xFF},
			},
		},
		{
			name: "Lexes 256-color yellow text from cube",
			text: "\033[38;5;226m",
			want: styleNode{
				fg: color.RGBA{0xFF, 0xFF, 0, 0xFF},
			},
		},
		{
			name: "Lexes 256-color cyan text from cube",
			text: "\033[38;5;51m",
			want: styleNode{
				fg: color.RGBA{0, 0xFF, 0xFF, 0xFF},
			},
		},
		{
			name: "Lexes 256-color magenta text from cube",
			text: "\033[38;5;201m",
			want: styleNode{
				fg: color.RGBA{0xFF, 0, 0xFF, 0xFF},
			},
		},
		{
			name: "Lexes 256-color white text from cube",
			text: "\033[38;5;231m",
			want: styleNode{
				fg: color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
			},
		},
		{
			name: "Lexes 256-color black text from grayscale",
			text: "\033[38;5;232m",
			want: styleNode{
				fg: color.RGBA{0, 0, 0, 0xFF},
			},
		},
		{
			name: "Lexes 256-color white text from grayscale",
			text: "\033[38;5;255m",
			want: styleNode{
				fg: color.RGBA{0xFF, 0xFF, 0xFF, 0xFF},
			},
		},
		{
			name: "Lexes true-color with no color channels",
			text: "\033[38;2m",
			want: styleNode{
				fg: color.RGBA{0, 0, 0, 0xFF},
			},
		},
		{
			name: "Lexes true-color with 1 color channel",
			text: "\033[38;2;1m",
			want: styleNode{
				fg: color.RGBA{1, 0, 0, 0xFF},
			},
		},
		{
			name: "Lexes true-color with 2 color channels",
			text: "\033[38;2;1;2m",
			want: styleNode{
				fg: color.RGBA{1, 2, 0, 0xFF},
			},
		},
		{
			name: "Lexes true-color foreground with 3 color channels",
			text: "\033[38;2;1;2;3m",
			want: styleNode{
				fg: color.RGBA{1, 2, 3, 0xFF},
			},
		},
		{
			name: "Lexes true-color background with 3 color channels",
			text: "\033[48;2;1;2;3m",
			want: styleNode{
				bg: color.RGBA{1, 2, 3, 0xFF},
			},
		},
	}

	for _, tt := range styleTests {
		t.Run(tt.name, func(t *testing.T) {
			nodes := lex(tt.text)
			next, stop := iter.Pull(nodes)
			defer stop()
			node, ok := next()
			if !ok {
				t.Fatalf("Expected a node")
			}
			styleNode, ok := node.(styleNode)
			if !ok {
				t.Fatalf("Expected a styleNode, got %T", node)
			}
			if want := tt.want.fg; want != nil {
				t.Logf("Checking styleNode.fg")
				assertColorsEqual(t, styleNode.fg, want)
			} else if styleNode.fg != nil {
				t.Errorf("styleNode.fg = %v, want nil", styleNode.fg)
			}

			if want := tt.want.bg; want != nil {
				t.Logf("Checking styleNode.bg")
				assertColorsEqual(t, styleNode.bg, want)
			} else if styleNode.bg != nil {
				t.Errorf("styleNode.bg = %v, want nil", styleNode.bg)
			}

			if want := tt.want.reset; styleNode.reset != want {
				t.Errorf("styleNode.reset = %t, want %t", styleNode.reset, want)
			}
		})
	}
}

func assertColorsEqual(t *testing.T, got, want color.Color) {
	t.Helper()
	if got == nil {
		t.Errorf("got nil, want color.Color")
		return
	}
	gr, gg, gb, ga := got.RGBA()
	wr, wg, wb, wa := want.RGBA()
	if gr != wr || gg != wg || gb != wb || ga != wa {
		t.Errorf("got %d %d %d %d, want %d %d %d %d", gr, gg, gb, ga, wr, wg, wb, wa)
	}
}
