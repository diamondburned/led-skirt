package animation

import (
	"math"

	"github.com/13rac1/fastmath"
	"libdb.so/led-skirt/internal/colors"
)

// BreathingFunction is a function that accepts a time from Tstart to Tend and
// returns the intensity of the LED at that time. The returned value should be
// between 0 and 0xFFFF inclusive.
type BreathingFunction func(d, max Milliseconds) uint16

// ShiftBreathingFunction returns a BreathingFunction that shifts the given
// BreathingFunction by half of the given duration. This causes the animation
// to start at a zero point instead of a peak.
func ShiftBreathingFunction(f BreathingFunction) BreathingFunction {
	return func(d, max Milliseconds) uint16 {
		// use max>>1 instead of max/2 for speed
		// the compiler probably optimizes this anyway
		return f((d+(max>>1))%max, max)
	}
}

// https://www.desmos.com/calculator/cwiccrnmx5

// BreatheSine returns the intensity of the LED at the given time using a sine
// function.
func BreatheSine(d, max Milliseconds) uint16 {
	return breatheSineFast(d, max)
}

func breatheSineAccurate(d, max Milliseconds) uint16 {
	t := CalculateDurationInU16(d, max)
	return uint16((0xFFFF/2)*math.Cos((math.Pi*float64(t))/(0xFFFF/2)) + (0xFFFF / 2))
}

func breatheSineFast(d, max Milliseconds) uint16 {
	t := CalculateDurationInU16(d, max)
	return uint16(fastmath.Cos16(t) + (0xFFFF / 2))
}

// BreatheLinear returns the intensity of the LED at the given time using a
// linear function.
func BreatheLinear(d, max Milliseconds) uint16 {
	const Tmid = 0xFFFF / 2
	t := CalculateDurationInU16(d, max)
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
// colors.RGB color.
func NewBreathingAnimation(rgb colors.RGB, breathingFunction BreathingFunction) AnimationFunc {
	if rgb.R == rgb.G && rgb.G == rgb.B {
		// fast path
		return func(d, max Milliseconds) colors.RGB {
			scale := breathingFunction(d, max)
			c := scaleColor(rgb.R, scale)
			return colors.RGB{R: c, G: c, B: c}
		}
	}
	return func(d, max Milliseconds) colors.RGB {
		scale := breathingFunction(d, max)
		return colors.RGB{
			R: scaleColor(rgb.R, scale),
			G: scaleColor(rgb.G, scale),
			B: scaleColor(rgb.B, scale),
		}
	}
}

func scaleColor(c uint8, scale uint16) uint8 {
	return uint8(uint32(c) * uint32(scale) / 0xFFFF)
}
