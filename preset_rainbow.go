//go:build preset_rainbow

package main

import (
	"time"

	"libdb.so/led-skirt/internal/animation"
	"libdb.so/led-skirt/internal/colors"
)

func init() {
	animationDuration = 30 * time.Second
	animationFunc = newRainbowAnimation()
}

func newRainbowAnimation() animation.AnimationFunc {
	hsv := colors.HSV{
		H: 0,
		S: 0xFF,
		V: colors.ScaleValue(0xFF, brightness),
	}

	return func(d, max animation.Milliseconds) colors.RGB {
		t := animation.CalculateDurationInU16(d, max)

		hsv := colors.HSV{
			H: colors.ScaleHue(t),
			S: hsv.S,
			V: hsv.V,
		}
		return hsv.ToRGB()
	}
}
