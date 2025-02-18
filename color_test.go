package main

import "testing"

func TestCliColor_UnmarshalText(t *testing.T) {
	tests := []struct {
		name  string
		text  string
		wantR uint8
		wantG uint8
		wantB uint8
		wantA uint8
	}{
		{
			name:  "6-digit hex",
			text:  "AABBCC",
			wantR: 0xAA,
			wantG: 0xBB,
			wantB: 0xCC,
			wantA: 0xFF,
		},
		{
			name:  "8-digit hex",
			text:  "AABBCCDD",
			wantR: 0xAA,
			wantG: 0xBB,
			wantB: 0xCC,
			wantA: 0xDD,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &cliColor{}
			if err := c.Set(tt.text); err != nil {
				t.Fatalf("UnmarshalText() error = %v", err)
			}
			if c.R != tt.wantR {
				t.Errorf("UnmarshalText() R = %v, want %v", c.R, tt.wantR)
			}
			if c.G != tt.wantG {
				t.Errorf("UnmarshalText() G = %v, want %v", c.G, tt.wantG)
			}
			if c.B != tt.wantB {
				t.Errorf("UnmarshalText() B = %v, want %v", c.B, tt.wantB)
			}
			if c.A != tt.wantA {
				t.Errorf("UnmarshalText() A = %v, want %v", c.A, tt.wantA)
			}
		})
	}
}
