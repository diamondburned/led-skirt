//go:build preset_rainbow

package main

import (
	"time"

	"libdb.so/led-skirt/internal/animation"
	"libdb.so/led-skirt/internal/colors"
	"libdb.so/led-skirt/internal/maths"
)

func init() {
	animationDuration = 30 * time.Second
	animationFunc = newRainbowAnimation()
}

func newRainbowAnimation() animation.AnimationFunc {
	hsv := colors.HSV{
		H: 0,
		S: 0xFF,
		V: colors.ScaleColor(0xFF, brightness),
	}

	return func(d, max animation.Milliseconds) colors.RGB {
		hsv := colors.HSV{
			H: maths.ScaleValue0(d, max, colors.MaxHue),
			S: hsv.S,
			V: hsv.V,
		}
		return hsv.ToRGB()
	}
}
