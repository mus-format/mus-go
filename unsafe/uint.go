package unsafe

import (
	"strconv"

	muscom "github.com/mus-format/mus-common-go"
	"github.com/mus-format/mus-go/raw"
)

func init() {
	setUpUintFuncs(strconv.IntSize)
}

var (
	marshalUint   func(v uint, bs []byte) (n int)
	unmarshalUint func(bs []byte) (v uint, n int, err error)
	sizeUint      func(v uint) int
	skipUint      func(bs []byte) (int, error)
)

// MarshalUint64 fills bs with the MUS encoding (Raw) of a uint64. Returns the
// number of used bytes.
func MarshalUint64(v uint64, bs []byte) (n int) {
	return marshalInteger64(v, bs)
}

// MarshalUint32 fills bs with the MUS encoding (Raw) of a uint32. Returns the
// number of used bytes.
func MarshalUint32(v uint32, bs []byte) (n int) {
	return marshalInteger32(v, bs)
}

// MarshalUint16 fills bs with the MUS encoding (Raw) of a uint16. Returns the
// number of used bytes.
func MarshalUint16(v uint16, bs []byte) (n int) {
	return marshalInteger16(v, bs)
}

// MarshalUint8 fills bs with the MUS encoding (Raw) of a uint8. Returns the
// number of used bytes.
func MarshalUint8(v uint8, bs []byte) (n int) {
	return marshalInteger8(v, bs)
}

// MarshalUint fills bs with the MUS encoding (Raw) of a uint. Returns the
// number of used bytes.
func MarshalUint(v uint, bs []byte) (n int) {
	return marshalUint(v, bs)
}

// -----------------------------------------------------------------------------
// UnmarshalUint64 parses a MUS-encoded (Raw) uint64 from bs. In addition to
// the uint64, it returns the number of used bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func UnmarshalUint64(bs []byte) (v uint64, n int, err error) {
	return unmarshalInteger64[uint64](bs)
}

// UnmarshalUint32 parses a MUS-encoded (Raw) uint32 from bs. In addition to
// the uint32, it returns the number of used bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func UnmarshalUint32(bs []byte) (v uint32, n int, err error) {
	return unmarshalInteger32[uint32](bs)
}

// UnmarshalUint16 parses a MUS-encoded (Raw) uint16 from bs. In addition to
// the uint16, it returns the number of used bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func UnmarshalUint16(bs []byte) (v uint16, n int, err error) {
	return unmarshalInteger16[uint16](bs)
}

// UnmarshalUint8 parses a MUS-encoded (Raw) uint8 from bs. In addition to
// the uint8, it returns the number of used bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func UnmarshalUint8(bs []byte) (v uint8, n int, err error) {
	return unmarshalInteger8[uint8](bs)
}

// UnmarshalUint parses a MUS-encoded (Raw) uint from bs. In addition to
// the uint, it returns the number of used bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func UnmarshalUint(bs []byte) (v uint, n int, err error) {
	return unmarshalUint(bs)
}

// -----------------------------------------------------------------------------
// SizeUint64 returns the size of a MUS-encoded (Raw) uint64.
func SizeUint64(v uint64) (n int) {
	return raw.SizeUint64(v)
}

// SizeUint32 returns the size of a MUS-encoded (Raw) uint32.
func SizeUint32(v uint32) (n int) {
	return raw.SizeUint32(v)
}

// SizeUint16 returns the size of a MUS-encoded (Raw) uint16.
func SizeUint16(v uint16) (n int) {
	return raw.SizeUint16(v)
}

// SizeUint8 returns the size of a MUS-encoded (Raw) uint8.
func SizeUint8(v uint8) (n int) {
	return raw.SizeUint8(v)
}

// SizeUint returns the size of a MUS-encoded (Raw) uint.
func SizeUint(v uint) (n int) {
	return sizeUint((v))
}

// -----------------------------------------------------------------------------
// SkipUint64 skips a MUS-encoded (Raw) uint64 in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func SkipUint64(bs []byte) (n int, err error) {
	return raw.SkipUint64(bs)
}

// SkipUint32 skips a MUS-encoded (Raw) uint32 in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func SkipUint32(bs []byte) (n int, err error) {
	return raw.SkipUint32(bs)
}

// SkipUint16 skips a MUS-encoded (Raw) uint16 in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func SkipUint16(bs []byte) (n int, err error) {
	return raw.SkipUint16(bs)
}

// SkipUint8 skips a MUS-encoded (Raw) uint8 in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func SkipUint8(bs []byte) (n int, err error) {
	return raw.SkipUint8(bs)
}

// SkipUint skips a MUS-encoded (Raw) uint in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be mus.ErrTooSmallByteSlice.
func SkipUint(bs []byte) (n int, err error) {
	return skipUint(bs)
}

func setUpUintFuncs(intSize int) {
	switch intSize {
	case 64:
		marshalUint = marshalInteger64[uint]
		unmarshalUint = unmarshalInteger64[uint]
	case 32:
		marshalUint = marshalInteger32[uint]
		unmarshalUint = unmarshalInteger32[uint]
	default:
		panic(muscom.ErrUnsupportedIntSize)
	}
	sizeUint = raw.SizeUint
	skipUint = raw.SkipUint
}
