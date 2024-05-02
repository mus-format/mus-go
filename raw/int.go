package raw

import (
	"strconv"

	com "github.com/mus-format/common-go"
)

func init() {
	setUpIntFuncs(strconv.IntSize)
}

var (
	marshalInt   func(v int, bs []byte) int
	unmarshalInt func(bs []byte) (int, int, error)
	sizeInt      int
	skipInt      func(bs []byte) (int, error)
)

// MarshalInt64 fills bs with the MUS encoding (Raw) of a int64 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalInt64(v int64, bs []byte) (n int) {
	return marshalInteger64(v, bs)
}

// MarshalInt32 fills bs with the MUS encoding (Raw) of a int32 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalInt32(v int32, bs []byte) (n int) {
	return marshalInteger32(v, bs)
}

// MarshalInt16 fills bs with the MUS encoding (Raw) of a int16 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalInt16(v int16, bs []byte) (n int) {
	return marshalInteger16(v, bs)
}

// MarshalInt8 fills bs with the MUS encoding (Raw) of a int8 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalInt8(v int8, bs []byte) (n int) {
	return marshalInteger8(v, bs)
}

// MarshalInt fills bs with the MUS encoding (Raw) of a int value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalInt(v int, bs []byte) (n int) {
	return marshalInt(v, bs)
}

// UnmarshalInt64 parses a MUS-encoded (Raw) int64 value from bs.
//
// In addition to the int64 value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice.
func UnmarshalInt64(bs []byte) (v int64, n int, err error) {
	return unmarshalInteger64[int64](bs)
}

// UnmarshalInt32 parses a MUS-encoded (Raw) int32 value from bs.
//
// In addition to the int32 value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice.
func UnmarshalInt32(bs []byte) (v int32, n int, err error) {
	return unmarshalInteger32[int32](bs)
}

// UnmarshalInt16 parses a MUS-encoded (Raw) int16 value from bs.
//
// In addition to the int16 value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice.
func UnmarshalInt16(bs []byte) (v int16, n int, err error) {
	return unmarshalInteger16[int16](bs)
}

// UnmarshalInt8 parses a MUS-encoded (Raw) int8 value from bs.
//
// In addition to the int8 value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice.
func UnmarshalInt8(bs []byte) (v int8, n int, err error) {
	return unmarshalInteger8[int8](bs)
}

// UnmarshalInt parses a MUS-encoded (Raw) int value from bs.
//
// In addition to the int value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice.
func UnmarshalInt(bs []byte) (v int, n int, err error) {
	return unmarshalInt(bs)
}

// SizeInt64 returns the size of a MUS-encoded (Raw) int64 value.
func SizeInt64(v int64) (n int) {
	return com.Num64RawSize
}

// SizeInt32 returns the size of a MUS-encoded (Raw) int32 value.
func SizeInt32(v int32) (n int) {
	return com.Num32RawSize
}

// SizeInt16 returns the size of a MUS-encoded (Raw) int16 value.
func SizeInt16(v int16) (n int) {
	return com.Num16RawSize
}

// SizeInt8 returns the size of a MUS-encoded (Raw) int8 value.
func SizeInt8(v int8) (n int) {
	return com.Num8RawSize
}

// SizeInt returns the size of a MUS-encoded (Raw) int value.
func SizeInt(v int) (n int) {
	return sizeInt
}

// SkipInt64 skips a MUS-encoded (Raw) int64 value.
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice.
func SkipInt64(bs []byte) (n int, err error) {
	return skipInteger64(bs)
}

// SkipInt32 skips a MUS-encoded (Raw) int32 value.
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice.
func SkipInt32(bs []byte) (n int, err error) {
	return skipInteger32(bs)
}

// SkipInt16 skips a MUS-encoded (Raw) int16 value.
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice.
func SkipInt16(bs []byte) (n int, err error) {
	return skipInteger16(bs)
}

// SkipInt8 skips a MUS-encoded (Raw) int8 value.
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice.
func SkipInt8(bs []byte) (n int, err error) {
	return skipInteger8(bs)
}

// SkipInt skips a MUS-encoded (Raw) int value.
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice.
func SkipInt(bs []byte) (n int, err error) {
	return skipInt(bs)
}

func setUpIntFuncs(intSize int) {
	switch intSize {
	case 64:
		marshalInt = marshalInteger64[int]
		unmarshalInt = unmarshalInteger64[int]
		sizeInt = com.Num64RawSize
		skipInt = skipInteger64
	case 32:
		marshalInt = marshalInteger32[int]
		unmarshalInt = unmarshalInteger32[int]
		sizeInt = com.Num32RawSize
		skipInt = skipInteger32
	default:
		panic(com.ErrUnsupportedIntSize)
	}
}
