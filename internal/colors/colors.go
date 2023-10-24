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
		R: ScaleValue(c.R, scale),
		G: ScaleValue(c.G, scale),
		B: ScaleValue(c.B, scale),
	}
}

// ScaleValue returns a new uint8 value with the brightness scaled by the given
// scale factor. It ensures that the resulting value is within the range of
// 0-255.
func ScaleValue(v uint8, scale float32) uint8 {
	scaled := float32(v) * scale
	if scaled > 255 {
		return 255
	}
	if scaled < 0 {
		return 0
	}
	return uint8(scaled)
}
