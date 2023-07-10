package unsafe

import (
	"strconv"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go/raw"
)

func init() {
	setUpIntFuncs(strconv.IntSize)
}

var (
	marshalInt   func(v int, bs []byte) (n int)
	unmarshalInt func(bs []byte) (v int, n int, err error)
	sizeInt      func(v int) int
	skipInt      func(bs []byte) (int, error)
)

// MarshalInt64 fills bs with the MUS encoding (Raw) of a int64. Returns the
// number of used bytes.
//
// It will panic if receives too small bs.
func MarshalInt64(v int64, bs []byte) (n int) {
	return marshalInteger64(v, bs)
}

// MarshalInt32 fills bs with the MUS encoding (Raw) of a int32. Returns the
// number of used bytes.
//
// It will panic if receives too small bs.
func MarshalInt32(v int32, bs []byte) (n int) {
	return marshalInteger32(v, bs)
}

// MarshalInt16 fills bs with the MUS encoding (Raw) of a int16. Returns the
// number of used bytes.
//
// It will panic if receives too small bs.
func MarshalInt16(v int16, bs []byte) (n int) {
	return marshalInteger16(v, bs)
}

// MarshalInt8 fills bs with the MUS encoding (Raw) of a int8. Returns the
// number of used bytes.
//
// It will panic if receives too small bs.
func MarshalInt8(v int8, bs []byte) (n int) {
	return marshalInteger8(v, bs)
}

// MarshalInt fills bs with the MUS encoding (Raw) of a int. Returns the
// number of used bytes.
//
// It will panic if receives too small bs.
func MarshalInt(v int, bs []byte) (n int) {
	return marshalInt(v, bs)
}

// -----------------------------------------------------------------------------
// UnmarshalInt64 parses a MUS-encoded (Raw) int64 from bs. In addition to
// the int64, it returns the number of used bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func UnmarshalInt64(bs []byte) (v int64, n int, err error) {
	return unmarshalInteger64[int64](bs)
}

// UnmarshalInt32 parses a MUS-encoded (Raw) int32 from bs. In addition to
// the int32, it returns the number of used bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func UnmarshalInt32(bs []byte) (v int32, n int, err error) {
	return unmarshalInteger32[int32](bs)
}

// UnmarshalInt16 parses a MUS-encoded (Raw) int16 from bs. In addition to
// the int16, it returns the number of used bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func UnmarshalInt16(bs []byte) (v int16, n int, err error) {
	return unmarshalInteger16[int16](bs)
}

// UnmarshalInt8 parses a MUS-encoded (Raw) int8 from bs. In addition to
// the int8, it returns the number of used bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func UnmarshalInt8(bs []byte) (v int8, n int, err error) {
	return unmarshalInteger8[int8](bs)
}

// UnmarshalInt parses a MUS-encoded (Raw) int from bs. In addition to
// the int, it returns the number of used bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func UnmarshalInt(bs []byte) (v int, n int, err error) {
	return unmarshalInt(bs)
}

// -----------------------------------------------------------------------------
// SizeInt64 returns the size of a MUS-encoded (Raw) int64.
func SizeInt64(v int64) (n int) {
	return raw.SizeInt64(v)
}

// SizeInt32 returns the size of a MUS-encoded (Raw) int32.
func SizeInt32(v int32) (n int) {
	return raw.SizeInt32(v)
}

// SizeInt16 returns the size of a MUS-encoded (Raw) int16.
func SizeInt16(v int16) (n int) {
	return raw.SizeInt16(v)
}

// SizeInt8 returns the size of a MUS-encoded (Raw) int8.
func SizeInt8(v int8) (n int) {
	return raw.SizeInt8(v)
}

// SizeInt returns the size of a MUS-encoded (Raw) int.
func SizeInt(v int) (n int) {
	return sizeInt(v)
}

// -----------------------------------------------------------------------------
// SkipInt64 skips a MUS-encoded (Raw) int64 in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func SkipInt64(bs []byte) (n int, err error) {
	return raw.SkipInt64(bs)
}

// SkipInt32 skips a MUS-encoded (Raw) int32 in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func SkipInt32(bs []byte) (n int, err error) {
	return raw.SkipInt32(bs)
}

// SkipInt16 skips a MUS-encoded (Raw) int16 in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func SkipInt16(bs []byte) (n int, err error) {
	return raw.SkipInt16(bs)
}

// SkipInt8 skips a MUS-encoded (Raw) int8 in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func SkipInt8(bs []byte) (n int, err error) {
	return raw.SkipInt8(bs)
}

// SkipInt skips a MUS-encoded (Raw) int in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func SkipInt(bs []byte) (n int, err error) {
	return skipInt(bs)
}

func setUpIntFuncs(intSize int) {
	switch intSize {
	case 64:
		marshalInt = marshalInteger64[int]
		unmarshalInt = unmarshalInteger64[int]
	case 32:
		marshalInt = marshalInteger32[int]
		unmarshalInt = unmarshalInteger32[int]
	default:
		panic(com.ErrUnsupportedIntSize)
	}
	sizeInt = raw.SizeInt
	skipInt = raw.SkipInt
}
