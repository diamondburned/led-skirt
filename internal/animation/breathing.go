package animation

import (
	"math"

	"github.com/13rac1/fastmath"
)

// BreathingFunction is a function that accepts a time from Tstart to Tend and
// returns the intensity of the LED at that time. The returned value should be
// between 0 and 0xFFFF inclusive.
type BreathingFunction func(d, max Milliseconds) uint16

// https://www.desmos.com/calculator/cwiccrnmx5

func scaleDuration(d, max Milliseconds) uint16 {
	return uint16(d * 0xFFFF / max)
}

// BreatheSine returns the intensity of the LED at the given time using a sine
// function.
func BreatheSine(d, max Milliseconds) uint16 {
	return breatheSineFast(d, max)
}

func breatheSineAccurate(d, max Milliseconds) uint16 {
	t := scaleDuration(d, max)
	return uint16((0xFFFF/2)*math.Cos((math.Pi*float64(t))/(0xFFFF/2)) + (0xFFFF / 2))
}

func breatheSineFast(d, max Milliseconds) uint16 {
	t := scaleDuration(d, max)
	return uint16(fastmath.Cos16(t) + (0xFFFF / 2))
}

// BreatheLinear returns the intensity of the LED at the given time using a
// linear function.
func BreatheLinear(d, max Milliseconds) uint16 {
	const Tmid = 0xFFFF / 2
	t := scaleDuration(d, max)
	v := 1 - (2 * (int32(t) - Tmid))
	if v < 0 {
		v = -v
	}
	return uint16(v)
}

var _ BreathingFunction = BreatheSine
var _ BreathingFunction = BreatheLinear

// NewBreathingAnimation returns an AnimationFunc that accepts a time from
// Tstart to Tend and returns the color of the LED at that time using the given
// breathing function and calculating the intensity of the LED using the given
// RGB color.
func NewBreathingAnimation(rgb RGB, breathingFunction BreathingFunction) AnimationFunc {
	if rgb.R == rgb.G && rgb.G == rgb.B {
		// fast path
		return func(d, max Milliseconds) RGB {
			scale := breathingFunction(d, max)
			c := uint8(uint16(rgb.R) * scale / 0xFFFF)
			return RGB{R: c, G: c, B: c}
		}
	}
	return func(d, max Milliseconds) RGB {
		scale := breathingFunction(d, max)
		return RGB{
			R: uint8(uint16(rgb.R) * scale / 0xFFFF),
			G: uint8(uint16(rgb.G) * scale / 0xFFFF),
			B: uint8(uint16(rgb.B) * scale / 0xFFFF),
		}
	}
}
