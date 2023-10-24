package main

import (
	"machine"
	"time"

	"libdb.so/led-skirt/internal/animation"
	"libdb.so/led-skirt/internal/colors"
	"tinygo.org/x/drivers/ws2812"
)

const brightness = 0.25 // operate at 25% brightness
const frameRate = 30    // operate at 30 FPS
const numLEDs = 30      // 30 LEDs in the skirt

func main() {
	// machine.LED.Configure(machine.PinConfig{Mode: machine.PinOutput})

	machine.D0.Configure(machine.PinConfig{Mode: machine.PinOutput})
	ledStrip := ws2812.New(machine.D0)

	// Convert the animation duration to milliseconds for performance.
	animationDurationMs := animation.DurationToMs(animationDuration)

	start := time.Now()
	for t := range time.Tick(time.Second / frameRate) {
		ms := animation.DurationToMs(t.Sub(start)) % animationDurationMs
		color := animationFunc(ms, animationDurationMs)
		drawColor(ledStrip, color)
	}
}

func drawColor(led ws2812.Device, c colors.RGB) {
	for i := 0; i < numLEDs; i++ {
		led.WriteByte(c.G) // WS2812 LEDs are GRB
		led.WriteByte(c.R)
		led.WriteByte(c.B)
	}
}
