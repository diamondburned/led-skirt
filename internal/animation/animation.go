package animation

import (
	"time"
)

// Milliseconds is the duration of the animation in milliseconds.
type Milliseconds uint32

// DurationToMs converts a standard library duration to milliseconds.
func DurationToMs(d time.Duration) Milliseconds {
	return Milliseconds(d / time.Millisecond)
}

// RGB represents 3 red, green, blue color channels in 8-bit unsigned integers.
type RGB struct {
	R, G, B uint8
}

// AnimationFunc is a function that accepts a time from Tstart to Tend and
// returns the color of the LED at that time.
type AnimationFunc func(d, max Milliseconds) RGB
