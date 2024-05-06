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
	length := len(v)
	n = varint.MarshalInt(length, bs)
	bs = bs[n:]
	if len(bs) < length {
		panic(mus.ErrTooSmallByteSlice)
	}
	return n + copy(bs, unsafe_mod.Slice(unsafe_mod.StringData(v), length))
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
// The lenVl argument specifies the string length Validator. If it returns
// an error and skip == true UnmarshalValidString skips the remaining bytes of
// the string.
//
// In addition to the string value, returns the number of used bytes and one of
// the mus.ErrTooSmallByteSlice, com.ErrOverflow, com.ErrNegativeLength or
// Validator errors.
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
	return unsafe_mod.String(&bs[n], length), l, nil
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
