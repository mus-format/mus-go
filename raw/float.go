package raw

import (
	"math"
)

// MarshalFloat64 fills bs with the MUS encoding (Raw) of a float64. Returns
// the number of used bytes.
//
// It will panic if receives too small bs.
func MarshalFloat64(v float64, bs []byte) (n int) {
	return marshalInteger64(math.Float64bits(v), bs)
}

// MarshalFloat32 fills bs with the MUS encoding (Raw) of a float32. Returns
// the number of used bytes.
//
// It will panic if receives too small bs.
func MarshalFloat32(v float32, bs []byte) (n int) {
	return marshalInteger32(math.Float32bits(v), bs)
}

// -----------------------------------------------------------------------------
// UnmarshalFloat64 parses a MUS-encoded (Raw) float64 from bs. In addition
// to the float64, it returns the number of used bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func UnmarshalFloat64(bs []byte) (v float64, n int, err error) {
	uv, n, err := unmarshalInteger64[uint64](bs)
	if err != nil {
		return
	}
	return math.Float64frombits(uv), n, nil
}

// UnmarshalFloat32 parses a MUS-encoded (Raw) float32 from bs. In addition
// to the float32, it returns the number of used bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func UnmarshalFloat32(bs []byte) (v float32, n int, err error) {
	uv, n, err := unmarshalInteger32[uint32](bs)
	if err != nil {
		return
	}
	return math.Float32frombits(uv), n, nil
}

// -----------------------------------------------------------------------------
// SizeFloat64 returns the size of a MUS-encoded (Raw) float64.
func SizeFloat64(v float64) (n int) {
	return sizeNum64(v)
}

// SizeFloat32 returns the size of a MUS-encoded (Raw) float32.
func SizeFloat32(v float32) (n int) {
	return sizeNum32(v)
}

// -----------------------------------------------------------------------------
// SkipFloat64 skips a MUS-encoded (Raw) float64 in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func SkipFloat64(bs []byte) (n int, err error) {
	return skipInteger64(bs)
}

// SkipFloat32 skips a MUS-encoded (Raw) float32 in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func SkipFloat32(bs []byte) (n int, err error) {
	return skipInteger32(bs)
}
