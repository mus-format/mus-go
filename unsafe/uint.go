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
	// Uint64 is an uint64 serializer.
	Uint64 = uint64Ser{}
	// Uint32 is an uint32 serializer.
	Uint32 = uint32Ser{}
	// Uint16 is an uint16 serializer.
	Uint16 = uint16Ser{}
	// Uint8 is an uint8 serializer.
	Uint8 = uint8Ser{}
	// Uint is an uint serializer.
	Uint = uintSer{}
)

var (
	marshalUint   func(v uint, bs []byte) int
	unmarshalUint func(bs []byte) (uint, int, error)
	sizeUint      int
	skipUint      func(bs []byte) (int, error)
)

type uint64Ser struct{}

// Marshal fills bs with an encoded (Raw) uint64 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s uint64Ser) Marshal(v uint64, bs []byte) (n int) {
	return marshalInteger64(v, bs)
}

// Unmarshal parses an encoded (Raw) uint64 value from bs.
//
// In addition to the uint64 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s uint64Ser) Unmarshal(bs []byte) (v uint64, n int, err error) {
	return unmarshalInteger64[uint64](bs)
}

// Size returns the size of an encoded (Raw) uint64 value.
func (s uint64Ser) Size(v uint64) (size int) {
	return com.Num64RawSize
}

// Skip skips an encoded (Raw) uint64 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func (s uint64Ser) Skip(bs []byte) (n int, err error) {
	return raw.SkipInteger64(bs)
}

// -----------------------------------------------------------------------------

type uint32Ser struct{}

// Marshal fills bs with an encoded (Raw) uint32 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s uint32Ser) Marshal(v uint32, bs []byte) (n int) {
	return marshalInteger32(v, bs)
}

// Unmarshal parses an encoded (Raw) uint32 value from bs.
//
// In addition to the uint32 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s uint32Ser) Unmarshal(bs []byte) (v uint32, n int, err error) {
	return unmarshalInteger32[uint32](bs)
}

// Size returns the size of an encoded (Raw) uint32 value.
func (s uint32Ser) Size(v uint32) (size int) {
	return com.Num32RawSize
}

// Skip skips an encoded (Raw) uint32 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func (s uint32Ser) Skip(bs []byte) (n int, err error) {
	return raw.SkipInteger32(bs)
}

// -----------------------------------------------------------------------------

type uint16Ser struct{}

// Marshal fills bs with an encoded (Raw) uint16 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s uint16Ser) Marshal(v uint16, bs []byte) (n int) {
	return marshalInteger16(v, bs)
}

// Unmarshal parses an encoded (Raw) uint16 value from bs.
//
// In addition to the uint16 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s uint16Ser) Unmarshal(bs []byte) (v uint16, n int, err error) {
	return unmarshalInteger16[uint16](bs)
}

// Size returns the size of an encoded (Raw) uint16 value.
func (s uint16Ser) Size(v uint16) (size int) {
	return com.Num16RawSize
}

// Skip skips an encoded (Raw) uint16 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func (s uint16Ser) Skip(bs []byte) (n int, err error) {
	return raw.SkipInteger16(bs)
}

// -----------------------------------------------------------------------------

type uint8Ser struct{}

// Marshal fills bs with an encoded (Raw) uint8 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s uint8Ser) Marshal(v uint8, bs []byte) (n int) {
	return marshalInteger8(v, bs)
}

// Unmarshal parses an encoded (Raw) uint8 value from bs.
//
// In addition to the uint8 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s uint8Ser) Unmarshal(bs []byte) (v uint8, n int, err error) {
	return unmarshalInteger8[uint8](bs)
}

// Size returns the size of an encoded (Raw) uint8 value.
func (s uint8Ser) Size(v uint8) (size int) {
	return com.Num8RawSize
}

// Skip skips an encoded (Raw) uint8 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func (s uint8Ser) Skip(bs []byte) (n int, err error) {
	return raw.SkipInteger8(bs)
}

// -----------------------------------------------------------------------------

type uintSer struct{}

// Marshal fills bs with an encoded (Raw) uint value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s uintSer) Marshal(v uint, bs []byte) (n int) {
	return marshalUint(v, bs)
}

// Unmarshal parses an encoded (Raw) uint value from bs.
//
// In addition to the uint value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s uintSer) Unmarshal(bs []byte) (v uint, n int, err error) {
	return unmarshalUint(bs)
}

// Size returns the size of an encoded (Raw) uint value.
func (s uintSer) Size(v uint) (size int) {
	return sizeUint
}

// Skip skips an encoded (Raw) uint value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func (s uintSer) Skip(bs []byte) (n int, err error) {
	return skipUint(bs)
}

// -----------------------------------------------------------------------------

func setUpUintFuncs(intSize int) {
	switch intSize {
	case 64:
		marshalUint = marshalInteger64[uint]
		unmarshalUint = unmarshalInteger64[uint]
		sizeUint = com.Num64RawSize
		skipUint = raw.SkipInteger64
	case 32:
		marshalUint = marshalInteger32[uint]
		unmarshalUint = unmarshalInteger32[uint]
		sizeUint = com.Num32RawSize
		skipUint = raw.SkipInteger32
	default:
		panic(com.ErrUnsupportedIntSize)
	}
}
