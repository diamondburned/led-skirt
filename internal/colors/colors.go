package colors

// RGB represents 3 red, green, blue color channels in 8-bit unsigned integers.
type RGB struct {
	R, G, B uint8
}

// ScaleBrightness returns a new RGB color with the brightness of each channel
// scaled by the given scale factor. It ensures that the resulting color is
// within the range of 0-255.
func ScaleBrightness(c RGB, scale float32) RGB {
	return RGB{
		R: ScaleColor(c.R, scale),
		G: ScaleColor(c.G, scale),
		B: ScaleColor(c.B, scale),
	}
}

// ScaleColor returns a new uint8 color value with the brightness scaled by the
// given scale factor. It ensures that the resulting value is within the range
// of 0-255.
func ScaleColor(v uint8, scale float32) uint8 {
	scaled := float32(v) * scale
	if scaled > 255 {
		return 255
	}
	if scaled < 0 {
		return 0
	}
	return uint8(scaled)
}
