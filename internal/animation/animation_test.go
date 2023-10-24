package animation

import (
	"testing"

	"libdb.so/led-skirt/internal/colors"
)

func TestInterpColor(t *testing.T) {
	type step struct {
		d uint16
		c colors.RGB
	}

	tests := []struct {
		name string
		c1   colors.RGB
		c2   colors.RGB
		want []step
	}{
		{
			name: "red to blue",
			c1:   colors.RGB{R: 255, G: 0, B: 0},
			c2:   colors.RGB{R: 0, G: 0, B: 255},
			want: []step{
				{d: 0x0000, c: colors.RGB{R: 255, G: 0, B: 0}},
				{d: 0x8000, c: colors.RGB{R: 127, G: 0, B: 127}},
				{d: 0xFFFF, c: colors.RGB{R: 0, G: 0, B: 255}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, step := range tt.want {
				got := linterpColor(tt.c1, tt.c2, step.d)
				if got != step.c {
					t.Errorf("linterpColor(%v, %v, %v) = %v, want %v", tt.c1, tt.c2, step.d, got, step.c)
				}
			}
		})
	}
}
