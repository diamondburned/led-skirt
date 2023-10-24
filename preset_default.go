package main

import (
	"time"

	"libdb.so/led-skirt/internal/animation"
	"libdb.so/led-skirt/internal/colors"
)

var animationDuration = 5 * time.Second
var animationFunc animation.AnimationFunc = animation.NewBreathingAnimation(
	colors.ScaleBrightness(colors.RGB{R: 255, G: 255, B: 255}, brightness),
	animation.BreatheSine,
)
