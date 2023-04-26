package ord

import (
	muscom "github.com/mus-format/mus-common-go"
	"github.com/mus-format/mus-go"
)

// MarshalBool fills bs with the MUS encoding of a bool. Returns the number of
// used bytes.
func MarshalBool(v bool, bs []byte) int {
	if v {
		bs[0] = 1
	} else {
		bs[0] = 0
	}
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
	if bs[0] == 0 {
		return false, 1, nil
	}
	if bs[0] == 1 {
		return true, 1, nil
	}
	return false, 0, muscom.ErrWrongFormat
}

// SizeBool returns the size of a MUS-encoded bool.
func SizeBool(v bool) (n int) {
	return 1
}

// SkipBool skips a MUS-encoded bool in bs. Returns the number of skiped bytes
// and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, or muscom.ErrWrongFormat.
func SkipBool(bs []byte) (n int, err error) {
	if len(bs) < 1 {
		return 0, mus.ErrTooSmallByteSlice
	}
	if bs[0] == 0 || bs[0] == 1 {
		return 1, nil
	}
	return 0, muscom.ErrWrongFormat
}
