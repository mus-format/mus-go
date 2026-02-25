package varint

import (
	"golang.org/x/exp/constraints"
)

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
	Int = IntSer{}
)

// int64 -----------------------------------------------------------------------

type int64Ser struct{}

// Marshal fills bs with an encoded (Varint) int64 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s int64Ser) Marshal(v int64, bs []byte) (n int) {
	return marshalUint(uint64(EncodeZigZag(v)), bs)
}

// Unmarshal parses an encoded (Varint) int64 value from bs.
//
// In addition to the int64 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s int64Ser) Unmarshal(bs []byte) (v int64, n int, err error) {
	uv, n, err := Uint64.Unmarshal(bs)
	if err != nil {
		return
	}
	return int64(DecodeZigZag(uv)), n, nil
}

// Size returns the size of an encoded (Varint) int64 value.
func (s int64Ser) Size(v int64) (size int) {
	return sizeUint(uint64(EncodeZigZag(v)))
}

// Skip skips the next encoded (Varint) int64 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s int64Ser) Skip(bs []byte) (n int, err error) {
	return Uint64.Skip(bs)
}

// int32 -----------------------------------------------------------------------

type int32Ser struct{}

// Marshal fills bs with an encoded (Varint) int32 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s int32Ser) Marshal(v int32, bs []byte) (n int) {
	return marshalUint(uint32(EncodeZigZag(v)), bs)
}

// Unmarshal parses an encoded (Varint) int32 value from bs.
//
// In addition to the int32 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s int32Ser) Unmarshal(bs []byte) (v int32, n int, err error) {
	uv, n, err := Uint32.Unmarshal(bs)
	if err != nil {
		return
	}
	return int32(DecodeZigZag(uv)), n, nil
}

// Size returns the size of an encoded (Varint) int32 value.
func (s int32Ser) Size(v int32) (size int) {
	return sizeUint(uint32(EncodeZigZag(v)))
}

// Skip skips the next encoded (Varint) int32 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s int32Ser) Skip(bs []byte) (n int, err error) {
	return Uint32.Skip(bs)
}

// int16 -----------------------------------------------------------------------

type int16Ser struct{}

// Marshal fills bs with an encoded (Varint) int16 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s int16Ser) Marshal(v int16, bs []byte) (n int) {
	return marshalUint(uint16(EncodeZigZag(v)), bs)
}

// Unmarshal parses an encoded (Varint) int16 value from bs.
//
// In addition to the int16 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s int16Ser) Unmarshal(bs []byte) (v int16, n int, err error) {
	uv, n, err := Uint16.Unmarshal(bs)
	if err != nil {
		return
	}
	return int16(DecodeZigZag(uv)), n, nil
}

// Size returns the size of an encoded (Varint) int16 value.
func (s int16Ser) Size(v int16) (size int) {
	return sizeUint(uint16(EncodeZigZag(v)))
}

// Skip skips the next encoded (Varint) int16 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s int16Ser) Skip(bs []byte) (n int, err error) {
	return Uint16.Skip(bs)
}

// int8 ------------------------------------------------------------------------

type int8Ser struct{}

// Marshal fills bs with an encoded (Varint) int8 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s int8Ser) Marshal(v int8, bs []byte) (n int) {
	return marshalUint(uint8(EncodeZigZag(v)), bs)
}

// Unmarshal parses an encoded (Varint) int8 value from bs.
//
// In addition to the int8 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s int8Ser) Unmarshal(bs []byte) (v int8, n int, err error) {
	uv, n, err := Uint8.Unmarshal(bs)
	if err != nil {
		return
	}
	return int8(DecodeZigZag(uv)), n, nil
}

// Size returns the size of an encoded (Varint) int8 value.
func (s int8Ser) Size(v int8) (size int) {
	return sizeUint(uint8(EncodeZigZag(v)))
}

// Skip skips the next encoded (Varint) int8 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s int8Ser) Skip(bs []byte) (n int, err error) {
	return Uint8.Skip(bs)
}

// int -------------------------------------------------------------------------

type IntSer struct{}

// Marshal fills bs with an encoded (Varint) int value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s IntSer) Marshal(v int, bs []byte) (n int) {
	return marshalUint(uint(EncodeZigZag(v)), bs)
}

// Unmarshal parses an encoded (Varint) int value from bs.
//
// In addition to the int value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s IntSer) Unmarshal(bs []byte) (v int, n int, err error) {
	uv, n, err := Uint.Unmarshal(bs)
	if err != nil {
		return
	}
	return int(DecodeZigZag(uv)), n, nil
}

// Size returns the size of an encoded (Varint) int value.
func (s IntSer) Size(v int) (size int) {
	return sizeUint(uint(EncodeZigZag(v)))
}

// Skip skips the next encoded (Varint) int value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s IntSer) Skip(bs []byte) (n int, err error) {
	return Uint.Skip(bs)
}

// -----------------------------------------------------------------------------

func EncodeZigZag[T constraints.Signed](t T) T {
	if t < 0 {
		return ^(t << 1)
	}
	return t << 1
}

func DecodeZigZag[T constraints.Unsigned](t T) T {
	if t&1 == 1 {
		return ^(t >> 1)
	}
	return t >> 1
}
