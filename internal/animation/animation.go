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
