package unsafe

import (
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/ord"
)

// MarshalBool fills bs with the MUS encoding of a bool. Returns the number of
// used bytes.
//
// It will panic if receives too small bs.
func MarshalBool(v bool, bs []byte) (n int) {
	*(*bool)(unsafe_mod.Pointer(&bs[0])) = v
	return 1
}

// UnmarshalBool parses a MUS-encoded bool from bs. In addition to the bool,
// it returns the number of used bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or com.ErrWrongFormat.
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

// SizeBool returns the size of a MUS-encoded bool.
func SizeBool(v bool) (n int) {
	return ord.SizeBool(v)
}

// SkipBool skips a MUS-encoded bool in bs. Returns the number of skiped bytes
// and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or com.ErrWrongFormat.
func SkipBool(bs []byte) (n int, err error) {
	return ord.SkipBool(bs)
}
