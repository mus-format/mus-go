package varint

import com "github.com/mus-format/common-go"

// MarshalByte fills bs with the MUS encoding (Varint) of a byte value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalByte(v byte, bs []byte) (n int) {
	return marshalUint(v, bs)
}

// UnmarshalByte parses a MUS-encoded (Varint) byte value from bs.
//
// In addition to the byte value, returns the number of used bytes and one of
// the mus.ErrTooSmallByteSlice or com.ErrOverflow errors.
func UnmarshalByte(bs []byte) (v byte, n int, err error) {
	return unmarshalUint[byte](com.Uint8MaxVarintLen, com.Uint8MaxLastByte,
		bs)
}

// SizeByte returns the size of a MUS-encoded (Varint) byte value.
func SizeByte(v byte) (size int) {
	return sizeUint(v)
}

// SkipByte skips a MUS-encoded (Varint) byte.
//
// Returns the number of skiped bytes and one of the mus.ErrTooSmallByteSlice or
// com.ErrOverflow errors.
func SkipByte(bs []byte) (n int, err error) {
	return skipUint(com.Uint8MaxVarintLen, com.Uint8MaxLastByte, bs)
}
