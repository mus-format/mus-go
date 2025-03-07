package unsafe

import (
	"math"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go/raw"
)

var (
	// Float64 is a float64 serializer.
	Float64 = float64Ser{}
	// Float32 is a float32 serializer.
	Float32 = float32Ser{}
)

type float64Ser struct{}

// Marshal fills bs with an encoded (Raw) float64 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s float64Ser) Marshal(v float64, bs []byte) (n int) {
	return marshalInteger64(math.Float64bits(v), bs)
}

// Unmarshal parses an encoded (Raw) float64 value from bs.
//
// In addition to the float64 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s float64Ser) Unmarshal(bs []byte) (v float64, n int, err error) {
	uv, n, err := unmarshalInteger64[uint64](bs)
	if err != nil {
		return
	}
	return math.Float64frombits(uv), n, nil
}

// Size returns the size of an encoded (Raw) float64 value.
func (s float64Ser) Size(v float64) (size int) {
	return com.Num64RawSize
}

// Skip skips an encoded (Raw) float64 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func (s float64Ser) Skip(bs []byte) (n int, err error) {
	return raw.SkipInteger64(bs)
}

// -----------------------------------------------------------------------------

type float32Ser struct{}

// Marshal fills bs with an encoded (Raw) float32 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s float32Ser) Marshal(v float32, bs []byte) (n int) {
	return marshalInteger32(math.Float32bits(v), bs)
}

// Unmarshal parses an encoded (Raw) float32 value from bs.
//
// In addition to the float32 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s float32Ser) Unmarshal(bs []byte) (v float32, n int, err error) {
	uv, n, err := unmarshalInteger32[uint32](bs)
	if err != nil {
		return
	}
	return math.Float32frombits(uv), n, nil
}

// Size returns the size of an encoded (Raw) float32 value.
func (s float32Ser) Size(v float32) (size int) {
	return com.Num32RawSize
}

// Skip skips an encoded (Raw) float32 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func (s float32Ser) Skip(bs []byte) (n int, err error) {
	return raw.SkipInteger32(bs)
}
