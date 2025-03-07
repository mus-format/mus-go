package varint

import (
	com "github.com/mus-format/common-go"
)

// Byte is a byte serializer.
var Byte = byteSer{}

type byteSer struct{}

// Marshal fills bs with an encoded (Varint) byte value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s byteSer) Marshal(v byte, bs []byte) (n int) {
	return marshalUint(v, bs)
}

// Unmarshal parses an encoded (Varint) byte value from bs.
//
// In addition to the byte value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow
func (s byteSer) Unmarshal(bs []byte) (v byte, n int, err error) {
	return unmarshalUint[byte](com.Uint8MaxVarintLen, com.Uint8MaxLastByte, bs)
}

// Size returns the size of an encoded (Varint) byte value.
func (s byteSer) Size(v byte) (size int) {
	return sizeUint(v)
}

// Skip skips an encoded (Varint) byte value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow
func (s byteSer) Skip(bs []byte) (n int, err error) {
	return skipUint(com.Uint8MaxVarintLen, com.Uint8MaxLastByte, bs)
}
