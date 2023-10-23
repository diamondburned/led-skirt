//go:build preset_trans

package main

import (
	"time"

	"libdb.so/led-skirt/internal/animation"
)

func init() {
	// animationFunc = newTransAnimation()
}

var transColors = []animation.RGB{
	{R: 85, G: 205, B: 252},
	{R: 255, G: 255, B: 255},
	{R: 247, G: 168, B: 184},
	{R: 85, G: 205, B: 252},
}

func newTransAnimation() animation.AnimationFunc {
	animations := make([]animation.AnimationFunc, len(transColors))
	for i, c := range transColors {
		animations[i] = animation.NewBreathingAnimation(c, animation.BreatheSine)
	}

	perAnimationDuration := animationDuration / time.Duration(len(animations))
	perAnimationMs := animation.DurationToMs(perAnimationDuration)

	return func(d, max animation.Milliseconds) animation.RGB {
		animationIndex := d / perAnimationMs
		animationMs := d % perAnimationMs
		return animations[animationIndex](animationMs, perAnimationMs)
	}
}
