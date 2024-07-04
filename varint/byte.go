package varint

import com "github.com/mus-format/common-go"

// MarshalByte fills bs with the encoding (Varint) of a byte value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalByte(v byte, bs []byte) (n int) {
	return marshalUint(v, bs)
}

// UnmarshalByte parses an encoded (Varint) byte value from bs.
//
// In addition to the byte value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice or com.ErrOverflow.
func UnmarshalByte(bs []byte) (v byte, n int, err error) {
	return unmarshalUint[byte](com.Uint8MaxVarintLen, com.Uint8MaxLastByte,
		bs)
}

// SizeByte returns the size of an encoded (Varint) byte value.
func SizeByte(v byte) (size int) {
	return sizeUint(v)
}

// SkipByte skips an encoded (Varint) byte.
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice or com.ErrOverflow.
func SkipByte(bs []byte) (n int, err error) {
	return skipUint(com.Uint8MaxVarintLen, com.Uint8MaxLastByte, bs)
}
