package unsafe

import (
	"strconv"

	com "github.com/mus-format/common-go"
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

// MarshalUint64 fills bs with an encoded (Raw) uint64 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalUint64(v uint64, bs []byte) (n int) {
	return marshalInteger64(v, bs)
}

// MarshalUint32 fills bs with an encoded (Raw) uint32 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalUint32(v uint32, bs []byte) (n int) {
	return marshalInteger32(v, bs)
}

// MarshalUint16 fills bs with an encoded (Raw) uint16 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalUint16(v uint16, bs []byte) (n int) {
	return marshalInteger16(v, bs)
}

// MarshalUint8 fills bs with an encoded (Raw) uint8 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalUint8(v uint8, bs []byte) (n int) {
	return marshalInteger8(v, bs)
}

// MarshalUint fills bs with an encoded (Raw) uint value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalUint(v uint, bs []byte) (n int) {
	return marshalUint(v, bs)
}

// UnmarshalUint64 parses an encoded (Raw) uint64 value from bs.
//
// In addition to the uint64 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func UnmarshalUint64(bs []byte) (v uint64, n int, err error) {
	return unmarshalInteger64[uint64](bs)
}

// UnmarshalUint32 parses an encoded (Raw) uint32 value from bs.
//
// In addition to the uint32 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func UnmarshalUint32(bs []byte) (v uint32, n int, err error) {
	return unmarshalInteger32[uint32](bs)
}

// UnmarshalUint16 parses an encoded (Raw) uint16 value from bs.
//
// In addition to the uint16 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func UnmarshalUint16(bs []byte) (v uint16, n int, err error) {
	return unmarshalInteger16[uint16](bs)
}

// UnmarshalUint8 parses an encoded (Raw) uint8 value from bs.
//
// In addition to the uint8 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func UnmarshalUint8(bs []byte) (v uint8, n int, err error) {
	return unmarshalInteger8[uint8](bs)
}

// UnmarshalUint parses an encoded (Raw) uint value from bs.
//
// In addition to the uint value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func UnmarshalUint(bs []byte) (v uint, n int, err error) {
	return unmarshalUint(bs)
}

// SizeUint64 returns the size of an encoded (Raw) uint64 value.
func SizeUint64(v uint64) (n int) {
	return raw.SizeUint64(v)
}

// SizeUint32 returns the size of an encoded (Raw) uint32 value.
func SizeUint32(v uint32) (n int) {
	return raw.SizeUint32(v)
}

// SizeUint16 returns the size of an encoded (Raw) uint16 value.
func SizeUint16(v uint16) (n int) {
	return raw.SizeUint16(v)
}

// SizeUint8 returns the size of an encoded (Raw) uint8 value.
func SizeUint8(v uint8) (n int) {
	return raw.SizeUint8(v)
}

// SizeUint returns the size of an encoded (Raw) uint value.
func SizeUint(v uint) (n int) {
	return sizeUint((v))
}

// SkipUint64 skips an encoded (Raw) uint64.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func SkipUint64(bs []byte) (n int, err error) {
	return raw.SkipUint64(bs)
}

// SkipUint32 skips an encoded (Raw) uint32.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func SkipUint32(bs []byte) (n int, err error) {
	return raw.SkipUint32(bs)
}

// SkipUint16 skips an encoded (Raw) uint16.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func SkipUint16(bs []byte) (n int, err error) {
	return raw.SkipUint16(bs)
}

// SkipUint8 skips an encoded (Raw) uint8.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func SkipUint8(bs []byte) (n int, err error) {
	return raw.SkipUint8(bs)
}

// SkipUint skips an encoded (Raw) uint.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
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
		panic(com.ErrUnsupportedIntSize)
	}
	sizeUint = raw.SizeUint
	skipUint = raw.SkipUint
}
