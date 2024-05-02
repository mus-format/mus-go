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
// In addition to the bool value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice or com.ErrWrongFormat.
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
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice or com.ErrWrongFormat.
func SkipBool(bs []byte) (n int, err error) {
	if len(bs) < 1 {
		return 0, mus.ErrTooSmallByteSlice
	}
	if bs[0] == 0 || bs[0] == 1 {
		return 1, nil
	}
	return 0, com.ErrWrongFormat
}
