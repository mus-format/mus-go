package varint

import "math"

// MarshalFloat64 fills bs with an encoded (Varint) float64 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalFloat64(v float64, bs []byte) int {
	return MarshalUint64(math.Float64bits(v), bs)
}

// MarshalFloat32 fills bs with an encoded (Varint) float32 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalFloat32(v float32, bs []byte) int {
	return MarshalUint32(math.Float32bits(v), bs)
}

// UnmarshalFloat64 parses an encoded (Varint) float64 value from bs.
//
// In addition to the float64 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func UnmarshalFloat64(bs []byte) (v float64, n int, err error) {
	uv, n, err := UnmarshalUint64(bs)
	if err != nil {
		return
	}
	v = math.Float64frombits(uv)
	return
}

// UnmarshalFloat32 parses an encoded (Varint) float32 value from bs.
//
// In addition to the float32 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func UnmarshalFloat32(bs []byte) (v float32, n int, err error) {
	uv, n, err := UnmarshalUint32(bs)
	if err != nil {
		return
	}
	v = math.Float32frombits(uv)
	return
}

// SizeFloat64 returns the size of an encoded (Varint) float64 value.
func SizeFloat64(v float64) (size int) {
	return SizeUint64(math.Float64bits(v))
}

// SizeFloat32 returns the size of an encoded (Varint) float32 value.
func SizeFloat32(v float32) (size int) {
	return SizeUint32(math.Float32bits(v))
}

// SkipFloat64 skips an encoded (Varint) float64 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func SkipFloat64(bs []byte) (n int, err error) {
	return SkipUint64(bs)
}

// SkipFloat32 skips an encoded (Varint) float32 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func SkipFloat32(bs []byte) (n int, err error) {
	return SkipUint32(bs)
}
