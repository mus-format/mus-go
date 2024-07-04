package varint

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"golang.org/x/exp/constraints"
)

// MarshalUint64 fills bs with the encoding (Varint) of a uint64 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalUint64(v uint64, bs []byte) (n int) {
	return marshalUint(v, bs)
}

// MarshalUint32 fills bs with the encoding (Varint) of a uint32 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalUint32(v uint32, bs []byte) (n int) {
	return marshalUint(v, bs)
}

// MarshalUint16 fills bs with the encoding (Varint) of a uint16 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalUint16(v uint16, bs []byte) (n int) {
	return marshalUint(v, bs)
}

// MarshalUint8 fills bs with the encoding (Varint) of a uint8 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalUint8(v uint8, bs []byte) (n int) {
	return marshalUint(v, bs)
}

// MarshalUint fills bs with the encoding (Varint) of a uint value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalUint(v uint, bs []byte) (n int) {
	return marshalUint(v, bs)
}

// UnmarshalUint64 parses an encoded (Varint) uint64 value from bs.
//
// In addition to the uint64 value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func UnmarshalUint64(bs []byte) (v uint64, n int, err error) {
	return unmarshalUint[uint64](com.Uint64MaxVarintLen,
		com.Uint64MaxLastByte,
		bs)
}

// UnmarshalUint32 parses an encoded (Varint) uint32 value from bs.
//
// In addition to the uint32 value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func UnmarshalUint32(bs []byte) (v uint32, n int, err error) {
	return unmarshalUint[uint32](com.Uint32MaxVarintLen,
		com.Uint32MaxLastByte,
		bs)
}

// UnmarshalUint16 parses an encoded (Varint) uint16 value from bs.
//
// In addition to the uint16 value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func UnmarshalUint16(bs []byte) (v uint16, n int, err error) {
	return unmarshalUint[uint16](com.Uint16MaxVarintLen,
		com.Uint16MaxLastByte,
		bs)
}

// UnmarshalUint8 parses an encoded (Varint) uint8 value from bs.
//
// In addition to the uint8 value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func UnmarshalUint8(bs []byte) (v uint8, n int, err error) {
	return unmarshalUint[uint8](com.Uint8MaxVarintLen,
		com.Uint8MaxLastByte,
		bs)
}

// UnmarshalUint parses an encoded (Varint) uint value from bs.
//
// In addition to the uint value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func UnmarshalUint(bs []byte) (v uint, n int, err error) {
	return unmarshalUint[uint](com.UintMaxVarintLen(),
		com.UintMaxLastByte(),
		bs)
}

// SizeUint64 returns the size of an encoded uint64 value.
func SizeUint64(v uint64) (size int) {
	return sizeUint(v)
}

// SizeUint32 returns the size of an encoded uint32 value.
func SizeUint32(v uint32) (size int) {
	return sizeUint(v)
}

// SizeUint16 returns the size of an encoded uint16 value.
func SizeUint16(v uint16) (size int) {
	return sizeUint(v)
}

// SizeUint8 returns the size of an encoded uint8 value.
func SizeUint8(v uint8) (size int) {
	return sizeUint(v)
}

// SizeUint returns the size of an encoded uint value.
func SizeUint(v uint) (size int) {
	return sizeUint(v)
}

// SkipUint64 skips an encoded uint64.
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func SkipUint64(bs []byte) (n int, err error) {
	return skipUint(com.Uint64MaxVarintLen, com.Uint64MaxLastByte, bs)
}

// SkipUint32 skips an encoded uint32.
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func SkipUint32(bs []byte) (n int, err error) {
	return skipUint(com.Uint32MaxVarintLen, com.Uint32MaxLastByte, bs)
}

// SkipUint16 skips an encoded uint16.
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func SkipUint16(bs []byte) (n int, err error) {
	return skipUint(com.Uint16MaxVarintLen, com.Uint16MaxLastByte, bs)
}

// SkipUint8 skips an encoded uint8.
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func SkipUint8(bs []byte) (n int, err error) {
	return skipUint(com.Uint8MaxVarintLen, com.Uint8MaxLastByte, bs)
}

// SkipUint skips an encoded uint.
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func SkipUint(bs []byte) (n int, err error) {
	return skipUint(com.UintMaxVarintLen(), com.UintMaxLastByte(), bs)
}

func marshalUint[T constraints.Unsigned](t T, bs []byte) (n int) {
	for t >= 0x80 {
		bs[n] = byte(t) | 0x80
		t >>= 7
		n++
	}
	bs[n] = byte(t)
	return n + 1
}

func unmarshalUint[T constraints.Unsigned](maxVarintLen int, maxLastByte byte,
	bs []byte) (t T, n int, err error) {
	if len(bs) == 0 {
		err = mus.ErrTooSmallByteSlice
		return
	}
	if bs[0] < 0x80 {
		return T(bs[0]), 1, nil
	}
	var (
		b     byte
		shift int
	)
	for n, b = range bs {
		n++
		if n == maxVarintLen && b > maxLastByte {
			return 0, n, com.ErrOverflow
		}
		if b < 0x80 {
			t = t | T(b)<<shift
			return
		}
		t = t | T(b&0x7F)<<shift
		shift += 7
	}
	return 0, n, mus.ErrTooSmallByteSlice
}

func sizeUint[T constraints.Unsigned](t T) (size int) {
	for t >= 0x80 {
		t >>= 7
		size++
	}
	return size + 1
}

func skipUint(maxVarintLen int, maxLastByte byte, bs []byte) (n int,
	err error) {
	if len(bs) == 0 {
		err = mus.ErrTooSmallByteSlice
		return
	}
	var b byte
	for n, b = range bs {
		if n == maxVarintLen && b > maxLastByte {
			return n, com.ErrOverflow
		}
		if b < 0x80 {
			n++
			return
		}
	}
	return n + 1, mus.ErrTooSmallByteSlice
}
