package unsafe

import (
	"github.com/mus-format/mus-go/raw"
)

// MarshalByte fills bs with the MUS encoding (Raw) of a byte value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalByte(v byte, bs []byte) (n int) {
	return marshalInteger8(v, bs)
}

// UnmarshalByte parses a MUS-encoded (Raw) byte value from bs.
//
// In addition to the byte value, returns the number of used bytes and the
// mus.ErrTooSmallByteSlice error.
func UnmarshalByte(bs []byte) (v byte, n int, err error) {
	return unmarshalInteger8[byte](bs)
}

// SizeByte returns the size of a MUS-encoded (Raw) byte value.
func SizeByte(v byte) (n int) {
	return raw.SizeByte(v)
}

// SkipByte skips a MUS-encoded (Raw) byte value.
//
// Returns the number of skiped bytes and the mus.ErrTooSmallByteSlice error.
func SkipByte(bs []byte) (n int, err error) {
	return raw.SkipByte(bs)
}
