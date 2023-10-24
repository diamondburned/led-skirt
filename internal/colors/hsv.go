package colors

// HSV is a color in the HSV color space.
type HSV struct {
	H uint16
	S uint8
	V uint8
}

// MaxHue defines the maximum value of the hue channel.
const MaxHue = hsvHueSteps - 1

// ScaleHue scales the hue value from a uint16 range [0x0000, 0xFFFF] to a
// uint16 range [0, MaxHue].
func ScaleHue(v uint16) uint16 {
	return uint16(uint32(v) * MaxHue / 0xFFFF)
}

// ToRGB converts the HSV color to RGB.
func (hsv HSV) ToRGB() RGB {
	var rgb RGB
	hsv2rgb(hsv, &rgb.R, &rgb.G, &rgb.B)
	return rgb
}

// The following code is adapted from the following link:
// https://www.vagrearg.org/content/hsvrgb
//
// It is licensed under the MIT license,
// Copyright (c) 2016  B. Stultiens.

const hsvHueSextant = 256
const hsvHueSteps = 6 * hsvHueSextant

func hsv2rgb(hsv HSV, r, g, b *uint8) {
	h := hsv.H
	s := hsv.S
	v := hsv.V

	if s == 0 {
		*r, *g, *b = v, v, v
		return
	}

	sextant := uint8(h >> 8)
	if sextant > 5 {
		panic("hsv2rgb: invalid hue value")
	}

	if sextant&2 != 0 {
		r, b = b, r
	}
	if sextant&4 != 0 {
		g, b = b, g
	}
	if sextant&6 == 0 {
		if sextant&1 == 0 {
			r, g = g, r
		}
	} else {
		if sextant&1 != 0 {
			r, g = g, r
		}
	}

	*g = v // Top level

	// Perform actual calculations

	// Bottom level: v * (1.0 - s)
	// --> (v * (255 - s) + error_corr + 1) / 256
	ww := uint16(v) * (255 - uint16(s)) // We don't use ^s to prevent size-promotion side effects
	ww += 1                             // Error correction
	ww += ww >> 8                       // Error correction
	*b = uint8(ww >> 8)

	hFraction := h & 0xff // 0...255

	if sextant&1 == 0 {
		// *r = ...slope_up...;
		d := uint32(v) * (uint32(255<<8) - uint32(uint16(s)*(256-hFraction)))
		d += d >> 8    // Error correction
		d += uint32(v) // Error correction
		*r = uint8(d >> 16)
	} else {
		// *r = ...slope_down...;
		d := uint32(v) * (uint32(255<<8) - uint32(uint16(s)*hFraction))
		d += d >> 8    // Error correction
		d += uint32(v) // Error correction
		*r = uint8(d >> 16)
	}
}
