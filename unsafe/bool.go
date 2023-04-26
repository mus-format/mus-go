package unsafe

import (
	unsafe_mod "unsafe"

	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/ord"
)

// MarshalBool fills bs with the MUS encoding of a bool. Returns the number of
// used bytes.
func MarshalBool(v bool, bs []byte) (n int) {
	*(*bool)(unsafe_mod.Pointer(&bs[0])) = v
	return 1
}

// UnmarshalBool parses a MUS-encoded bool from bs. In addition to the bool,
// it returns the number of used bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrWrongFormat.
func UnmarshalBool(bs []byte) (v bool, n int, err error) {
	if len(bs) < 1 {
		return false, 0, mus.ErrTooSmallByteSlice
	}
	return *(*bool)(unsafe_mod.Pointer(&bs[0])), 1, nil
}

// SizeBool returns the size of a MUS-encoded bool.
func SizeBool(v bool) (n int) {
	return ord.SizeBool(v)
}

// SkipBool skips a MUS-encoded bool in bs. Returns the number of skiped bytes
// and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrWrongFormat.
func SkipBool(bs []byte) (n int, err error) {
	return ord.SkipBool(bs)
}
