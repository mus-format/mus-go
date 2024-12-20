package unsafe

import (
	"math"

	"github.com/mus-format/mus-go/raw"
)

// MarshalFloat64 fills bs with an encoded (Raw) float64 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalFloat64(v float64, bs []byte) (n int) {
	return marshalInteger64(math.Float64bits(v), bs)
}

// MarshalFloat32 fills bs with an encoded (Raw) float32 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalFloat32(v float32, bs []byte) (n int) {
	return marshalInteger32(math.Float32bits(v), bs)
}

// UnmarshalFloat64 parses an encoded (Raw) float64 value from bs.
//
// In addition to the float64 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func UnmarshalFloat64(bs []byte) (v float64, n int, err error) {
	uv, n, err := unmarshalInteger64[uint64](bs)
	if err != nil {
		return
	}
	return math.Float64frombits(uv), n, nil
}

// UnmarshalFloat32 parses an encoded (Raw) float32 value from bs.
//
// In addition to the float32 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func UnmarshalFloat32(bs []byte) (v float32, n int, err error) {
	uv, n, err := unmarshalInteger32[uint32](bs)
	if err != nil {
		return
	}
	return math.Float32frombits(uv), n, nil
}

// SizeFloat64 returns the size of an encoded (Raw) float64 value.
func SizeFloat64(v float64) (n int) {
	return raw.SizeFloat64(v)
}

// SizeFloat32 returns the size of an encoded (Raw) float32 value.
func SizeFloat32(v float32) (n int) {
	return raw.SizeFloat32(v)
}

// SkipFloat64 skips an encoded (Raw) float64 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func SkipFloat64(bs []byte) (n int, err error) {
	return raw.SkipFloat64(bs)
}

// SkipFloat32 skips an encoded (Raw) float32 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func SkipFloat32(bs []byte) (n int, err error) {
	return raw.SkipFloat32(bs)
}
