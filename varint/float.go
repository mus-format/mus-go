package varint

import "math"

// MarshalFloat64 fills bs with the MUS encoding (Varint) of a float64. Returns
// the number of used bytes.
func MarshalFloat64(v float64, bs []byte) int {
	return MarshalUint64(math.Float64bits(v), bs)
}

// MarshalFloat32 fills bs with the MUS encoding (Varint) of a float32. Returns
// the number of used bytes.
func MarshalFloat32(v float32, bs []byte) int {
	return MarshalUint32(math.Float32bits(v), bs)
}

// -----------------------------------------------------------------------------
// UnmarshalFloat64 parses a MUS-encoded (Varint) float64 from bs. In addition
// to the float64, it returns the number of used bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func UnmarshalFloat64(bs []byte) (v float64, n int, err error) {
	uv, n, err := UnmarshalUint64(bs)
	if err != nil {
		return
	}
	v = math.Float64frombits(uv)
	return
}

// UnmarshalFloat32 parses a MUS-encoded (Varint) float32 from bs. In addition
// to the float32, it returns the number of used bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func UnmarshalFloat32(bs []byte) (v float32, n int, err error) {
	uv, n, err := UnmarshalUint32(bs)
	if err != nil {
		return
	}
	v = math.Float32frombits(uv)
	return
}

// -----------------------------------------------------------------------------
// SizeFloat64 returns the size of a MUS-encoded (Varint) float64.
func SizeFloat64(v float64) (size int) {
	return SizeUint64(math.Float64bits(v))
}

// SizeFloat32 returns the size of a MUS-encoded (Varint) float32.
func SizeFloat32(v float32) (size int) {
	return SizeUint32(math.Float32bits(v))
}

// -----------------------------------------------------------------------------
// SkipFloat64 skips a MUS-encoded (Varint) float64 in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func SkipFloat64(bs []byte) (n int, err error) {
	return SkipUint64(bs)
}

// SkipFloat32 skips a MUS-encoded (Varint) float32 in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func SkipFloat32(bs []byte) (n int, err error) {
	return SkipUint32(bs)
}