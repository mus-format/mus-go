package raw

import (
	"strconv"

	com "github.com/mus-format/common-go"
)

func init() {
	setUpIntFuncs(strconv.IntSize)
}

var (
	// Int64 is an int64 serializer.
	Int64 = int64Ser{}
	// Int32 is an int32 serializer.
	Int32 = int32Ser{}
	// Int16 is an int16 serializer.
	Int16 = int16Ser{}
	// Int8 is an int8 serializer.
	Int8 = int8Ser{}
	// Int is an int serializer.
	Int = intSer{}
)

var (
	marshalInt   func(v int, bs []byte) int
	unmarshalInt func(bs []byte) (int, int, error)
	sizeInt      int
	skipInt      func(bs []byte) (int, error)
)

// -----------------------------------------------------------------------------

type int64Ser struct{}

// Marshal fills bs with an encoded (Raw) int64 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s int64Ser) Marshal(v int64, bs []byte) (n int) {
	return marshalInteger64[int64](v, bs)
}

// Unmarshal parses an encoded (Raw) int64 value from bs.
//
// In addition to the int64 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s int64Ser) Unmarshal(bs []byte) (v int64, n int, err error) {
	return unmarshalInteger64[int64](bs)
}

// Size returns the size of an encoded (Raw) int64 value.
func (s int64Ser) Size(v int64) (size int) {
	return com.Num64RawSize
}

// Skip skips an encoded (Raw) int64 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func (s int64Ser) Skip(bs []byte) (n int, err error) {
	return SkipInteger64(bs)
}

// -----------------------------------------------------------------------------

type int32Ser struct{}

// Marshal fills bs with an encoded (Raw) int32 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s int32Ser) Marshal(v int32, bs []byte) (n int) {
	return marshalInteger32[int32](v, bs)
}

// Unmarshal parses an encoded (Raw) int32 value from bs.
//
// In addition to the int32 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s int32Ser) Unmarshal(bs []byte) (v int32, n int, err error) {
	return unmarshalInteger32[int32](bs)
}

// Size returns the size of an encoded (Raw) int32 value.
func (s int32Ser) Size(v int32) (size int) {
	return com.Num32RawSize
}

// Skip skips an encoded (Raw) int32 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func (s int32Ser) Skip(bs []byte) (n int, err error) {
	return SkipInteger32(bs)
}

// -----------------------------------------------------------------------------

type int16Ser struct{}

// Marshal fills bs with an encoded (Raw) int16 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s int16Ser) Marshal(v int16, bs []byte) (n int) {
	return marshalInteger16[int16](v, bs)
}

// Unmarshal parses an encoded (Raw) int16 value from bs.
//
// In addition to the int16 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s int16Ser) Unmarshal(bs []byte) (v int16, n int, err error) {
	return unmarshalInteger16[int16](bs)
}

// Size returns the size of an encoded (Raw) int16 value.
func (s int16Ser) Size(v int16) (size int) {
	return com.Num16RawSize
}

// Skip skips an encoded (Raw) int16 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func (s int16Ser) Skip(bs []byte) (n int, err error) {
	return SkipInteger16(bs)
}

// -----------------------------------------------------------------------------

type int8Ser struct{}

// Marshal fills bs with an encoded (Raw) int8 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s int8Ser) Marshal(v int8, bs []byte) (n int) {
	return marshalInteger8[int8](v, bs)
}

// Unmarshal parses an encoded (Raw) int8 value from bs.
//
// In addition to the int8 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s int8Ser) Unmarshal(bs []byte) (v int8, n int, err error) {
	return unmarshalInteger8[int8](bs)
}

// Size returns the size of an encoded (Raw) int8 value.
func (s int8Ser) Size(v int8) (size int) {
	return com.Num8RawSize
}

// Skip skips an encoded (Raw) int8 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func (s int8Ser) Skip(bs []byte) (n int, err error) {
	return SkipInteger8(bs)
}

// -----------------------------------------------------------------------------

type intSer struct{}

// Marshal fills bs with an encoded (Raw) int value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s intSer) Marshal(v int, bs []byte) (n int) {
	return marshalInt(v, bs)
}

// Unmarshal parses an encoded (Raw) int value from bs.
//
// In addition to the int value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s intSer) Unmarshal(bs []byte) (v int, n int, err error) {
	return unmarshalInt(bs)
}

// Size returns the size of an encoded (Raw) int value.
func (s intSer) Size(v int) (size int) {
	return sizeInt
}

// Skip skips an encoded (Raw) int value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func (s intSer) Skip(bs []byte) (n int, err error) {
	return skipInt(bs)
}

func setUpIntFuncs(intSize int) {
	switch intSize {
	case 64:
		marshalInt = marshalInteger64[int]
		unmarshalInt = unmarshalInteger64[int]
		sizeInt = com.Num64RawSize
		skipInt = SkipInteger64
	case 32:
		marshalInt = marshalInteger32[int]
		unmarshalInt = unmarshalInteger32[int]
		sizeInt = com.Num32RawSize
		skipInt = SkipInteger32
	default:
		panic(com.ErrUnsupportedIntSize)
	}
}
