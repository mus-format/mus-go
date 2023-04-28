package varint

import (
	muscom "github.com/mus-format/mus-common-go"
)

// MarshalByte fills bs with the MUS encoding (Varint) of a byte. Returns the
// number of used bytes.
//
// It will panic if receives too small bs.
func MarshalByte(v byte, bs []byte) (n int) {
	return marshalUint(v, bs)
}

// UnmarshalByte parses a MUS-encoded (Varint) byte from bs. In addition to the
// byte, it returns the number of used bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func UnmarshalByte(bs []byte) (v byte, n int, err error) {
	return unmarshalUint[byte](muscom.Uint8MaxVarintLen, muscom.Uint8MaxLastByte,
		bs)
}

// SizeByte returns the size of a MUS-encoded (Varint) byte.
func SizeByte(v byte) (size int) {
	return sizeUint(v)
}

// SkipByte skips a MUS-encoded (Varint) byte in bs. Returns the number of
// skiped bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrOverflow.
func SkipByte(bs []byte) (n int, err error) {
	return skipUint(muscom.Uint8MaxVarintLen, muscom.Uint8MaxLastByte, bs)
}
