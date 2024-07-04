package varint

import (
	"golang.org/x/exp/constraints"
)

// MarshalInt64 fills bs with the encoding (Varint) of a int64 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalInt64(v int64, bs []byte) (n int) {
	return marshalUint(uint64(EncodeZigZag(v)), bs)
}

// MarshalInt32 fills bs with the encoding (Varint) of a int32 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalInt32(v int32, bs []byte) (n int) {
	return marshalUint(uint32(EncodeZigZag(v)), bs)
}

// MarshalInt16 fills bs with the encoding (Varint) of a int16 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalInt16(v int16, bs []byte) (n int) {
	return marshalUint(uint16(EncodeZigZag(v)), bs)
}

// MarshalInt8 fills bs with the encoding (Varint) of a int8 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalInt8(v int8, bs []byte) (n int) {
	return marshalUint(uint8(EncodeZigZag(v)), bs)
}

// MarshalInt fills bs with the encoding (Varint) of a int value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalInt(v int, bs []byte) (n int) {
	return marshalUint(uint(EncodeZigZag(v)), bs)
}

// UnmarshalInt64 parses an encoded (Varint) int64 value from bs.
//
// In addition to the int64 value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func UnmarshalInt64(bs []byte) (v int64, n int, err error) {
	uv, n, err := UnmarshalUint64(bs)
	if err != nil {
		return
	}
	return int64(DecodeZigZag(uv)), n, nil
}

// UnmarshalInt32 parses an encoded (Varint) int32 value from bs.
//
// In addition to the int32 value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func UnmarshalInt32(bs []byte) (v int32, n int, err error) {
	uv, n, err := UnmarshalUint32(bs)
	if err != nil {
		return
	}
	return int32(DecodeZigZag(uv)), n, nil
}

// UnmarshalInt16 parses an encoded (Varint) int16 value from bs.
//
// In addition to the int16 value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func UnmarshalInt16(bs []byte) (v int16, n int, err error) {
	uv, n, err := UnmarshalUint16(bs)
	if err != nil {
		return
	}
	return int16(DecodeZigZag(uv)), n, nil
}

// UnmarshalInt8 parses an encoded (Varint) int8 value from bs.
//
// In addition to the int8 value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func UnmarshalInt8(bs []byte) (v int8, n int, err error) {
	uv, n, err := UnmarshalUint8(bs)
	if err != nil {
		return
	}
	return int8(DecodeZigZag(uv)), n, nil
}

// UnmarshalInt parses an encoded (Varint) int value from bs.
//
// In addition to the int value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func UnmarshalInt(bs []byte) (v int, n int, err error) {
	uv, n, err := UnmarshalUint(bs)
	if err != nil {
		return
	}
	return int(DecodeZigZag(uv)), n, nil
}

// SizeInt64 returns the size of an encoded (Varint) int64 value.
func SizeInt64(v int64) int {
	return sizeUint(uint64(EncodeZigZag(v)))
}

// SizeInt32 returns the size of an encoded (Varint) int32 value.
func SizeInt32(v int32) int {
	return SizeUint32(uint32(EncodeZigZag(v)))
}

// SizeInt16 returns the size of an encoded (Varint) int16 value.
func SizeInt16(v int16) (size int) {
	return SizeUint16(uint16(EncodeZigZag(v)))
}

// SizeInt8 returns the size of an encoded (Varint) int8 value.
func SizeInt8(v int8) (size int) {
	return SizeUint8(uint8(EncodeZigZag(v)))
}

// SizeInt returns the size of an encoded (Varint) int value.
func SizeInt(v int) (size int) {
	return SizeUint(uint(EncodeZigZag(v)))
}

// SkipInt64 skips an encoded (Varint) int64 value.
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func SkipInt64(bs []byte) (n int, err error) {
	return SkipUint64(bs)
}

// SkipInt32 skips an encoded (Varint) int32 value.
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func SkipInt32(bs []byte) (n int, err error) {
	return SkipUint32(bs)
}

// SkipInt16 skips an encoded (Varint) int16 value.
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func SkipInt16(bs []byte) (n int, err error) {
	return SkipUint16(bs)
}

// SkipInt8 skips an encoded (Varint) int8 value.
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func SkipInt8(bs []byte) (n int, err error) {
	return SkipUint8(bs)
}

// SkipInt skips an encoded (Varint) int value.
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func SkipInt(bs []byte) (n int, err error) {
	return SkipUint(bs)
}

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
