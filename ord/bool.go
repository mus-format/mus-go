package ord

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
)

// MarshalBool fills bs with the MUS encoding of a bool value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalBool(v bool, bs []byte) int {
	if v {
		bs[0] = 1
	} else {
		bs[0] = 0
	}
	return 1
}

// UnmarshalBool parses a MUS-encoded bool value from bs.
//
// In addition to the bool value, returns the number of used bytes and one of
// the mus.ErrTooSmallByteSlice or com.ErrWrongFormat errors.
func UnmarshalBool(bs []byte) (v bool, n int, err error) {
	if len(bs) < 1 {
		return false, 0, mus.ErrTooSmallByteSlice
	}
	switch bs[0] {
	case 0:
		return false, 1, nil
	case 1:
		return true, 1, nil
	default:
		return false, 0, com.ErrWrongFormat
	}
}

// SizeBool returns the size of a MUS-encoded bool value.
func SizeBool(v bool) (n int) {
	return 1
}

// SkipBool skips a MUS-encoded bool value.
//
// Returns the number of skiped bytes and one of the mus.ErrTooSmallByteSlice
// or com.ErrWrongFormat errors.
func SkipBool(bs []byte) (n int, err error) {
	if len(bs) < 1 {
		return 0, mus.ErrTooSmallByteSlice
	}
	if bs[0] == 0 || bs[0] == 1 {
		return 1, nil
	}
	return 0, com.ErrWrongFormat
}
