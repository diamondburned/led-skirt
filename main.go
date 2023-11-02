package main

import (
	"machine"
	"time"

	"libdb.so/led-skirt/internal/animation"
	"libdb.so/led-skirt/internal/colors"
	"tinygo.org/x/drivers/ws2812"
)

const utilization = 1   // utilize 100% of the LEDs
const brightness = 0.15 // operate at 25% brightness
const frameRate = 30    // operate at 30 FPS
const numLEDs = 60      // 60 LEDs in the skirt

func main() {
	// machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})

	machine.D0.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledStrip := ws2812.New(machine.D0)

	// Convert the animation duration to milliseconds for performance.
	animationDurationMs := animation.DurationToMs(animationDuration)

	var skipEvery int
	if utilization < 1 {
		skipEvery = int(1 / utilization)
	}

	start := time.Now()
	for t := range time.Tick(time.Second / frameRate) {
		ms := animation.DurationToMs(t.Sub(start)) % animationDurationMs
		color := animationFunc(ms, animationDurationMs)
		drawColor(ledStrip, color, skipEvery)
	}
}

func drawColor(led ws2812.Device, c colors.RGB, skipEvery int) {
	for i := 0; i < numLEDs; i++ {
		c := c
		if skipEvery != 0 && i%skipEvery == 0 {
			c = colors.RGB{}
		}

		led.WriteByte(c.G) // WS2812 LEDs are GRB
		led.WriteByte(c.R)
		led.WriteByte(c.B)
	}
}
