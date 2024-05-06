package ord

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/varint"
)

// MarshalString fills bs with the MUS encoding of a string value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalString(v string, bs []byte) (n int) {
	length := len(v)
	n = varint.MarshalInt(length, bs)
	if len(bs) < n+length {
		panic(mus.ErrTooSmallByteSlice)
	}
	return n + copy(bs[n:], v)
}

// UnmarshalString parses a MUS-encoded string value from bs.
//
// In addition to the string value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice, com.ErrOverflow or com.ErrNegativeLength.
func UnmarshalString(bs []byte) (v string, n int, err error) {
	return UnmarshalValidString(nil, false, bs)
}

// UnmarshalValidString parses a MUS-encoded valid string value from bs.
//
// The lenVl argument specifies the string length Validator. If it returns
// an error and skip == true UnmarshalValidString skips the remaining bytes of
// the string.
//
// In addition to the string value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice, com.ErrOverflow, com.ErrNegativeLength or
// Validator error.
func UnmarshalValidString(lenVl com.Validator[int], skip bool, bs []byte) (
	v string, n int, err error) {
	length, n, err := varint.UnmarshalInt(bs)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	if lenVl != nil {
		if err = lenVl.Validate(length); err != nil {
			if skip {
				n += length
			}
			return
		}
	}
	if length == 0 {
		return
	}
	l := n + length
	if len(bs) < l {
		err = mus.ErrTooSmallByteSlice
		return
	}
	return string(bs[n:l]), l, nil
}

// SizeString returns the size of a MUS-encoded string value.
func SizeString(v string) (n int) {
	length := len(v)
	return varint.SizeInt(length) + length
}

// SkipString skips a MUS-encoded string value.
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice, com.ErrOverflow or mus.ErrNegativeLength.
func SkipString(bs []byte) (n int, err error) {
	length, n, err := varint.UnmarshalInt(bs)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	if len(bs[n:]) < length {
		err = mus.ErrTooSmallByteSlice
		return
	}
	return n + length, nil
}
