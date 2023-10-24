//go:build preset_trans

package main

import (
	"time"

	"libdb.so/led-skirt/internal/animation"
	"libdb.so/led-skirt/internal/colors"
)

func init() {
	animationDuration = 16 * time.Second
	animationFunc = newTransAnimation()
}

var transColors = []colors.RGB{
	{R: 10, G: 150, B: 204},
	{R: 219, G: 52, B: 116},
	{R: 255, G: 255, B: 255},
	{R: 219, G: 52, B: 116},
	// {R: 10, G: 150, B: 204},
}

func newTransAnimation() animation.AnimationFunc {
	animations := make([]animation.AnimationFunc, len(transColors))
	for i, c := range transColors {
		animations[i] = animation.NewBreathingAnimation(
			colors.ScaleBrightness(c, brightness),
			// Induce a shift in duration so that the animation starts at a zero
			// point instead of a peak.
			animation.ShiftBreathingFunction(animation.BreatheSine),
		)
	}

	perAnimationDuration := animationDuration / time.Duration(len(animations))
	perAnimationMs := animation.DurationToMs(perAnimationDuration)

	return func(d, max animation.Milliseconds) colors.RGB {
		animationIndex := d / perAnimationMs
		animationMs := d % perAnimationMs
		return animations[animationIndex](animationMs, perAnimationMs)
	}
}
