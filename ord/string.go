package ord

import (
	muscom "github.com/mus-format/mus-common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/varint"
)

// MarshalString fills bs with the MUS encoding of a string. Returns the number
// of used bytes.
//
// It will panic if receives too small bs.
func MarshalString(v string, bs []byte) (n int) {
	n = varint.MarshalInt(len(v), bs)
	if len(bs[n:]) < len(v) {
		panic(mus.ErrTooSmallByteSlice)
	}
	return n + copy(bs[n:], v)
}

// UnmarshalString parses a MUS-encoded string from bs. In addition to the
// string, it returns the number of used bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, muscom.ErrOverflow, or
// muscom.ErrNegativeLength.
func UnmarshalString(bs []byte) (v string, n int, err error) {
	return UnmarshalValidString(nil, false, bs)
}

// UnmarshalValidString parses a MUS-encoded valid string from bs. In addition
// to the string, it returns the number of used bytes and an error.
//
// The maxLength argument specifies the string length Validator. If it returns
// an error and skip == true UnmarshalValidString skips the remaining bytes of
// the string.
//
// The error returned by UnmarshalValidString can be one of
// mus.ErrTooSmallByteSlice, muscom.ErrOverflow, muscom.ErrNegativeLength, or a
// Validator error.
func UnmarshalValidString(maxLength muscom.Validator[int], skip bool, bs []byte) (
	v string, n int, err error) {
	length, n, err := varint.UnmarshalInt(bs)
	if err != nil || length == 0 {
		return
	}
	if length < 0 {
		err = muscom.ErrNegativeLength
		return
	}
	if len(bs[n:]) < length {
		err = mus.ErrTooSmallByteSlice
		return
	}
	if maxLength != nil {
		if err = maxLength.Validate(length); err != nil {
			if skip {
				n += length
			}
			return
		}
	}
	return string(bs[n : length+1]), n + length, nil
}

// SizeString returns the size of a MUS-encoded string.
func SizeString(v string) (n int) {
	return varint.SizeInt(len(v)) + len(v)
}

// SkipString skips a MUS-encoded string in bs. Returns the number of skiped
// bytes and an error.
//
// The error can be one of mus.ErrTooSmallByteSlice, muscom.ErrOverflow, or
// mus.ErrNegativeLength.
func SkipString(bs []byte) (n int, err error) {
	length, n, err := varint.UnmarshalInt(bs)
	if err != nil {
		return
	}
	if length < 0 {
		err = muscom.ErrNegativeLength
		return
	}
	if len(bs[n:]) < length {
		err = mus.ErrTooSmallByteSlice
		return
	}
	return n + length, nil
}
