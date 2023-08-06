package unsafe

import (
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/varint"
)

// MarshalString fills bs with the MUS encoding of a string value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalString(v string, bs []byte) (n int) {
	n = varint.MarshalInt(len(v), bs)
	if len(bs[n:]) < len(v) {
		panic(mus.ErrTooSmallByteSlice)
	}
	return n + copy(bs[n:], unsafe_mod.Slice(unsafe_mod.StringData(v), len(v)))
}

// UnmarshalString parses a MUS-encoded string value from bs.
//
// In addition to the string value, returns the number of used bytes and one of
// the mus.ErrTooSmallByteSlice, com.ErrOverflow or com.ErrNegativeLength
// errors.
func UnmarshalString(bs []byte) (v string, n int, err error) {
	return UnmarshalValidString(nil, false, bs)
}

// UnmarshalValidString parses a MUS-encoded valid string value from bs.
//
// The maxLength argument specifies the string length Validator. If it returns
// an error and skip == true UnmarshalValidString skips the remaining bytes of
// the string.
//
// In addition to the string value, returns the number of used bytes and one of
// the mus.ErrTooSmallByteSlice, com.ErrOverflow, com.ErrNegativeLength or
// Validator errors.
func UnmarshalValidString(maxLength com.Validator[int], skip bool, bs []byte) (
	v string, n int, err error) {
	length, n, err := varint.UnmarshalInt(bs)
	if err != nil || length == 0 {
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
	if maxLength != nil {
		if err = maxLength.Validate(length); err != nil {
			if skip {
				n += length
			}
			return
		}
	}
	c := bs[n : n+length]
	return unsafe_mod.String(&c[0], len(c)), n + length, nil
}

// SizeString returns the size of a MUS-encoded string value.
func SizeString(v string) (n int) {
	return ord.SizeString(v)
}

// SkipString skips a MUS-encoded string.
//
// Returns the number of skiped bytes and one of the mus.ErrTooSmallByteSlice,
// com.ErrOverflow or mus.ErrNegativeLength errors.
func SkipString(bs []byte) (n int, err error) {
	return ord.SkipString(bs)
}
