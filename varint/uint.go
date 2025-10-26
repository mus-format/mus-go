package varint

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"golang.org/x/exp/constraints"
)

var (
	// Uint64 is a uint64 serializer.
	Uint64 = uint64Ser{}
	// Uint32 is a uint32 serializer.
	Uint32 = uint32Ser{}
	// Uint16 is a uint16 serializer.
	Uint16 = uint16Ser{}
	// Uint8 is a uint8 serializer.
	Uint8 = uint8Ser{}
	// Uint is a uint serializer.
	Uint = uintSer{}
)

type uint64Ser struct{}

// Marshal fills bs with an encoded (Varint) uint64 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s uint64Ser) Marshal(v uint64, bs []byte) (n int) {
	return marshalUint(v, bs)
}

// Unmarshal parses an encoded (Varint) uint64 value from bs.
//
// In addition to the uint64 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s uint64Ser) Unmarshal(bs []byte) (v uint64, n int, err error) {
	return unmarshalUint[uint64](com.Uint64MaxVarintLen, com.Uint64MaxLastByte,
		bs)
}

// Size returns the size of an encoded (Varint) uint64 value.
func (s uint64Ser) Size(v uint64) (size int) {
	return sizeUint(v)
}

// Skip skips the next encoded (Varint) uint64 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s uint64Ser) Skip(bs []byte) (n int, err error) {
	return skipUint(com.Uint64MaxVarintLen, com.Uint64MaxLastByte, bs)
}

// -----------------------------------------------------------------------------

type uint32Ser struct{}

// Marshal fills bs with an encoded (Varint) uint32 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s uint32Ser) Marshal(v uint32, bs []byte) (n int) {
	return marshalUint(v, bs)
}

// Unmarshal parses an encoded (Varint) uint32 value from bs.
//
// In addition to the uint32 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s uint32Ser) Unmarshal(bs []byte) (v uint32, n int, err error) {
	return unmarshalUint[uint32](com.Uint32MaxVarintLen, com.Uint32MaxLastByte,
		bs)
}

// Size returns the size of an encoded (Varint) uint32 value.
func (s uint32Ser) Size(v uint32) (size int) {
	return sizeUint(v)
}

// Skip skips the next encoded (Varint) uint32 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s uint32Ser) Skip(bs []byte) (n int, err error) {
	return skipUint(com.Uint32MaxVarintLen, com.Uint32MaxLastByte, bs)
}

// -----------------------------------------------------------------------------

type uint16Ser struct{}

// Marshal fills bs with an encoded (Varint) uint16 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s uint16Ser) Marshal(v uint16, bs []byte) (n int) {
	return marshalUint(v, bs)
}

// Unmarshal parses an encoded (Varint) uint16 value from bs.
//
// In addition to the uint16 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s uint16Ser) Unmarshal(bs []byte) (v uint16, n int, err error) {
	return unmarshalUint[uint16](com.Uint16MaxVarintLen, com.Uint16MaxLastByte,
		bs)
}

// Size returns the size of an encoded (Varint) uint16 value.
func (s uint16Ser) Size(v uint16) (size int) {
	return sizeUint(v)
}

// Skip skips the next encoded (Varint) uint16 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s uint16Ser) Skip(bs []byte) (n int, err error) {
	return skipUint(com.Uint16MaxVarintLen, com.Uint16MaxLastByte, bs)
}

// -----------------------------------------------------------------------------

type uint8Ser struct{}

// Marshal fills bs with an encoded (Varint) uint8 value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s uint8Ser) Marshal(v uint8, bs []byte) (n int) {
	return marshalUint(v, bs)
}

// Unmarshal parses an encoded (Varint) uint8 value from bs.
//
// In addition to the uint8 value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s uint8Ser) Unmarshal(bs []byte) (v uint8, n int, err error) {
	return unmarshalUint[uint8](com.Uint8MaxVarintLen, com.Uint8MaxLastByte,
		bs)
}

// Size returns the size of an encoded (Varint) uint8 value.
func (s uint8Ser) Size(v uint8) (size int) {
	return sizeUint(v)
}

// Skip skips the next encoded (Varint) uint8 value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s uint8Ser) Skip(bs []byte) (n int, err error) {
	return skipUint(com.Uint8MaxVarintLen, com.Uint8MaxLastByte, bs)
}

// -----------------------------------------------------------------------------

type uintSer struct{}

// Marshal fills bs with an encoded (Varint) uint value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s uintSer) Marshal(v uint, bs []byte) (n int) {
	return marshalUint(v, bs)
}

// Unmarshal parses an encoded (Varint) uint value from bs.
//
// In addition to the uint value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s uintSer) Unmarshal(bs []byte) (v uint, n int, err error) {
	return unmarshalUint[uint](com.UintMaxVarintLen(), com.UintMaxLastByte(), bs)
}

// Size returns the size of an encoded (Varint) uint value.
func (s uintSer) Size(v uint) (size int) {
	return sizeUint(v)
}

// Skip skips the next encoded (Varint) uint value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func (s uintSer) Skip(bs []byte) (n int, err error) {
	return skipUint(com.UintMaxVarintLen(), com.UintMaxLastByte(), bs)
}

// -----------------------------------------------------------------------------

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
	bs []byte,
) (t T, n int, err error) {
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
	err error,
) {
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
