package raw

import (
	com "github.com/mus-format/common-go"
)

// Byte is a byte serializer.
var Byte = byteSer{}

type byteSer struct{}

// Marshal fills bs with an encoded (Raw) byte value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s byteSer) Marshal(v byte, bs []byte) (n int) {
	return marshalInteger8(v, bs)
}

// Unmarshal parses an encoded (Raw) byte value from bs.
//
// In addition to the byte value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice.
func (s byteSer) Unmarshal(bs []byte) (v byte, n int, err error) {
	return unmarshalInteger8[byte](bs)
}

// Size returns the size of an encoded (Raw) byte value.
func (s byteSer) Size(v byte) (size int) {
	return com.Num8RawSize
}

// Skip skips an encoded (Raw) byte value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func (s byteSer) Skip(bs []byte) (n int, err error) {
	return SkipInteger8(bs)
}
