package varint

import "math"

var (
	// Float64 is a float64 serializer.
	Float64 = float64Ser{}
	// Float32 is a float32 serializer.
	Float32 = float32Ser{}
)

type float64Ser struct{}

// Marshal fills bs with an encoded (Varint) float64 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s float64Ser) Marshal(v float64, bs []byte) (n int) {
	return marshalUint(math.Float64bits(v), bs)
}

// Unmarshal parses an encoded (Varint) float64 value from bs.
//
// In addition to the float64 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow
func (s float64Ser) Unmarshal(bs []byte) (v float64, n int, err error) {
	uv, n, err := Uint64.Unmarshal(bs)
	if err != nil {
		return
	}
	v = math.Float64frombits(uv)
	return
}

// Size returns the size of an encoded (Varint) float64 value.
func (s float64Ser) Size(v float64) (size int) {
	return sizeUint(math.Float64bits(v))
}

// Skip skips an encoded (Varint) float64 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow
func (s float64Ser) Skip(bs []byte) (n int, err error) {
	return Uint64.Skip(bs)
}

// -----------------------------------------------------------------------------

type float32Ser struct{}

// Marshal fills bs with an encoded (Varint) float32 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s float32Ser) Marshal(v float32, bs []byte) (n int) {
	return marshalUint(math.Float32bits(v), bs)
}

// Unmarshal parses an encoded (Varint) float32 value from bs.
//
// In addition to the float32 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow
func (s float32Ser) Unmarshal(bs []byte) (v float32, n int, err error) {
	uv, n, err := Uint32.Unmarshal(bs)
	if err != nil {
		return
	}
	v = math.Float32frombits(uv)
	return
}

// Size returns the size of an encoded (Varint) float32 value.
func (s float32Ser) Size(v float32) (size int) {
	return sizeUint(math.Float32bits(v))
}

// Skip skips an encoded (Varint) float32 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow
func (s float32Ser) Skip(bs []byte) (n int, err error) {
	return Uint32.Skip(bs)
}
