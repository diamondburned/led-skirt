package maths

// UnsignedInteger is an interface that represents an unsigned integer type.
// It does not support 64-bit integers.
type UnsignedInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32
}

// ScaleValue0 scales v from the range [0, maxv] to the range [0, maxout].
// It performs arithmetic in uint32 to avoid overflow.
func ScaleValue0[T1, T2 UnsignedInteger](v, maxv T1, maxout T2) T2 {
	return T2(uint32(v) * uint32(maxout) / uint32(maxv))
}
