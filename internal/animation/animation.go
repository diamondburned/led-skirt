package animation

import (
	"time"

	"libdb.so/led-skirt/internal/colors"
	"libdb.so/led-skirt/internal/maths"
)

// Milliseconds is the duration of the animation in milliseconds.
type Milliseconds uint32

// DurationToMs converts a standard library duration to milliseconds.
func DurationToMs(d time.Duration) Milliseconds {
	return Milliseconds(d / time.Millisecond)
}

// CalculateDurationInU16 calculates the duration scale for a given duration
// and maximum duration in units of 1/65536. It returns a value between 0 and
// 0xFFFF inclusive.
func CalculateDurationInU16(d, max Milliseconds) uint16 {
	return maths.ScaleValue0(d, max, uint16(0xFFFF))
}

// AnimationFunc is a function that accepts a time from Tstart to Tend and
// returns the color of the LED at that time.
type AnimationFunc func(d, max Milliseconds) colors.RGB

// NewLinearInterpolator returns an AnimationFunc that interpolates between
// the given colors using a linear function.
func NewLinearInterpolator(c1, c2 colors.RGB) AnimationFunc {
	return func(d, max Milliseconds) colors.RGB {
		return linterpColor(c1, c2, CalculateDurationInU16(d, max))
	}
}

func linterpColor(c1, c2 colors.RGB, scale uint16) colors.RGB {
	return colors.RGB{
		R: uint8(((uint32(c2.R) - uint32(c1.R)) * uint32(scale) / 0xFFFF) + uint32(c1.R)),
		G: uint8(((uint32(c2.G) - uint32(c1.G)) * uint32(scale) / 0xFFFF) + uint32(c1.G)),
		B: uint8(((uint32(c2.B) - uint32(c1.B)) * uint32(scale) / 0xFFFF) + uint32(c1.B)),
	}
}
