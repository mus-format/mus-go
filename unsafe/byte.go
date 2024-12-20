package unsafe

import (
	"github.com/mus-format/mus-go/raw"
)

// MarshalByte fills bs with an encoded (Raw) byte value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalByte(v byte, bs []byte) (n int) {
	return marshalInteger8(v, bs)
}

// UnmarshalByte parses an encoded (Raw) byte value from bs.
//
// In addition to the byte value and the number of used bytes, it may also return
// return mus.ErrTooSmallByteSlice.
func UnmarshalByte(bs []byte) (v byte, n int, err error) {
	return unmarshalInteger8[byte](bs)
}

// SizeByte returns the size of an encoded (Raw) byte value.
func SizeByte(v byte) (n int) {
	return raw.SizeByte(v)
}

// SkipByte skips an encoded (Raw) byte value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice.
func SkipByte(bs []byte) (n int, err error) {
	return raw.SkipByte(bs)
}
