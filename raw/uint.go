package raw

import (
	"strconv"

	com "github.com/mus-format/common-go"
)

func init() {
	setUpUintFuncs(strconv.IntSize)
}

var (
	marshalUint   func(v uint, bs []byte) int
	unmarshalUint func(bs []byte) (uint, int, error)
	sizeUint      func(v uint) int
	skipUint      func(bs []byte) (int, error)
)

// MarshalUint64 fills bs with the MUS encoding (Raw) of a uint64 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalUint64(v uint64, bs []byte) (n int) {
	return marshalInteger64(v, bs)
}

// MarshalUint32 fills bs with the MUS encoding (Raw) of a uint32 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalUint32(v uint32, bs []byte) (n int) {
	return marshalInteger32(v, bs)
}

// MarshalUint16 fills bs with the MUS encoding (Raw) of a uint16 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalUint16(v uint16, bs []byte) (n int) {
	return marshalInteger16(v, bs)
}

// MarshalUint8 fills bs with the MUS encoding (Raw) of a uint8 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalUint8(v uint8, bs []byte) (n int) {
	return marshalInteger8(v, bs)
}

// MarshalUint fills bs with the MUS encoding (Raw) of a uint value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalUint(v uint, bs []byte) (n int) {
	return marshalUint(v, bs)
}

// -----------------------------------------------------------------------------
// UnmarshalUint64 parses a MUS-encoded (Raw) uint64 value from bs.
//
// In addition to the uint64, returns the number of used bytes and the
// mus.ErrTooSmallByteSlice error.
func UnmarshalUint64(bs []byte) (v uint64, n int, err error) {
	return unmarshalInteger64[uint64](bs)
}

// UnmarshalUint32 parses a MUS-encoded (Raw) uint32 value from bs.
//
// In addition to the uint32, returns the number of used bytes and the
// mus.ErrTooSmallByteSlice error.
func UnmarshalUint32(bs []byte) (v uint32, n int, err error) {
	return unmarshalInteger32[uint32](bs)
}

// UnmarshalUint16 parses a MUS-encoded (Raw) uint16 value from bs.
//
// In addition to the uint16, returns the number of used bytes and the
// mus.ErrTooSmallByteSlice error.
func UnmarshalUint16(bs []byte) (v uint16, n int, err error) {
	return unmarshalInteger16[uint16](bs)
}

// UnmarshalUint8 parses a MUS-encoded (Raw) uint8 value from bs.
//
// In addition to the uint8, returns the number of used bytes and the
// mus.ErrTooSmallByteSlice error.
func UnmarshalUint8(bs []byte) (v uint8, n int, err error) {
	return unmarshalInteger8[uint8](bs)
}

// UnmarshalUint parses a MUS-encoded (Raw) uint value from bs.
//
// In addition to the uint, returns the number of used bytes and the
// mus.ErrTooSmallByteSlice error.
func UnmarshalUint(bs []byte) (v uint, n int, err error) {
	return unmarshalUint(bs)
}

// -----------------------------------------------------------------------------
// SizeUint64 returns the size of a MUS-encoded (Raw) uint64 value.
func SizeUint64(v uint64) (n int) {
	return sizeNum64(v)
}

// SizeUint32 returns the size of a MUS-encoded (Raw) uint32 value.
func SizeUint32(v uint32) (n int) {
	return sizeNum32(v)
}

// SizeUint16 returns the size of a MUS-encoded (Raw) uint16 value.
func SizeUint16(v uint16) (n int) {
	return sizeInteger16(v)
}

// SizeUint8 returns the size of a MUS-encoded (Raw) uint8 value.
func SizeUint8(v uint8) (n int) {
	return sizeInteger8(v)
}

// SizeUint returns the size of a MUS-encoded (Raw) uint value.
func SizeUint(v uint) (n int) {
	return sizeUint(v)
}

// -----------------------------------------------------------------------------
// SkipUint64 skips a MUS-encoded (Raw) uint64.
//
// Returns the number of skiped bytes and the mus.ErrTooSmallByteSlice error.
func SkipUint64(bs []byte) (n int, err error) {
	return skipInteger64(bs)
}

// SkipUint32 skips a MUS-encoded (Raw) uint32.
//
// Returns the number of skiped bytes and the mus.ErrTooSmallByteSlice error.
func SkipUint32(bs []byte) (n int, err error) {
	return skipInteger32(bs)
}

// SkipUint16 skips a MUS-encoded (Raw) uint16.
//
// Returns the number of skiped bytes and the mus.ErrTooSmallByteSlice error.
func SkipUint16(bs []byte) (n int, err error) {
	return skipInteger16(bs)
}

// SkipUint8 skips a MUS-encoded (Raw) uint8.

// Returns the number of skiped bytes and the mus.ErrTooSmallByteSlice error.
func SkipUint8(bs []byte) (n int, err error) {
	return skipInteger8(bs)
}

// SkipUint skips a MUS-encoded (Raw) uint.
//
// Returns the number of skiped bytes and the mus.ErrTooSmallByteSlice error.
func SkipUint(bs []byte) (n int, err error) {
	return skipUint(bs)
}

func setUpUintFuncs(intSize int) {
	switch intSize {
	case 64:
		marshalUint = marshalInteger64[uint]
		unmarshalUint = unmarshalInteger64[uint]
		sizeUint = sizeNum64[uint]
		skipUint = skipInteger64
	case 32:
		marshalUint = marshalInteger32[uint]
		unmarshalUint = unmarshalInteger32[uint]
		sizeUint = sizeNum32[uint]
		skipUint = skipInteger32
	default:
		panic(com.ErrUnsupportedIntSize)
	}
}
