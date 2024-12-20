package raw

import com "github.com/mus-format/common-go"

// MarshalByte fills bs with an encoded (Raw) byte value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalByte(v byte, bs []byte) (n int) {
	return marshalInteger8(v, bs)
}

// UnmarshalByte parses an encoded (Raw) byte value from bs.
//
// In addition to the byte value and the number of used bytes, it can
// return mus.ErrTooSmallByteSlice.
func UnmarshalByte(bs []byte) (v byte, n int, err error) {
	return unmarshalInteger8[byte](bs)
}

// SizeByte returns the size of an encoded (Raw) byte value.
func SizeByte(v byte) (n int) {
	return com.Num8RawSize
}

// SkipByte skips an encoded (Raw) byte value.
//
// In addition to the number of skipped bytes, it can return
// mus.ErrTooSmallByteSlice.
func SkipByte(bs []byte) (n int, err error) {
	return skipInteger8(bs)
}
