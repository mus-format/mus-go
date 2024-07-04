package unsafe

import (
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/ord"
)

// MarshalBool fills bs with the encoding of a bool value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalBool(v bool, bs []byte) (n int) {
	*(*bool)(unsafe_mod.Pointer(&bs[0])) = v
	return 1
}

// UnmarshalBool parses an encoded bool value from bs.
//
// In addition to the bool, returns the number of used bytes and one of the
// mus.ErrTooSmallByteSlice or com.ErrWrongFormat errors.
func UnmarshalBool(bs []byte) (v bool, n int, err error) {
	if len(bs) < 1 {
		return false, 0, mus.ErrTooSmallByteSlice
	}
	if bs[0] > 1 {
		err = com.ErrWrongFormat
		return
	}
	return *(*bool)(unsafe_mod.Pointer(&bs[0])), 1, nil
}

// SizeBool returns the size of an encoded bool value.
func SizeBool(v bool) (n int) {
	return ord.SizeBool(v)
}

// SkipBool skips an encoded bool value.
//
// Returns the number of skiped bytes and one of the mus.ErrTooSmallByteSlice or
// com.ErrWrongFormat errors.
func SkipBool(bs []byte) (n int, err error) {
	return ord.SkipBool(bs)
}
