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
// The lenM argument specifies the Marshaller for the string length.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalString(v string, lenM mus.Marshaller[int], bs []byte) (n int) {
	length := len(v)
	if lenM == nil {
		n = varint.MarshalPositiveInt(length, bs)
	} else {
		n = lenM.MarshalMUS(length, bs)
	}
	if len(bs) < n+length {
		panic(mus.ErrTooSmallByteSlice)
	}
	return n + copy(bs[n:], unsafe_mod.Slice(unsafe_mod.StringData(v), length))
}

// UnmarshalString parses a MUS-encoded string value from bs.
//
// The lenU argument specifies the Unmarshaller for the string length.
//
// In addition to the string value, returns the number of used bytes and one of
// the mus.ErrTooSmallByteSlice, com.ErrOverflow or com.ErrNegativeLength
// errors.
func UnmarshalString(lenU mus.Unmarshaller[int], bs []byte) (v string, n int,
	err error) {
	return UnmarshalValidString(lenU, nil, false, bs)
}

// UnmarshalValidString parses a MUS-encoded valid string value from bs.
//
// The lenU argument specifies the Unmarshaller for the string length.
// The lenVl argument specifies the string length Validator. If it returns
// an error and skip == true UnmarshalValidString skips the remaining bytes of
// the string.
//
// In addition to the string value, returns the number of used bytes and one of
// the mus.ErrTooSmallByteSlice, com.ErrOverflow, com.ErrNegativeLength or
// Validator errors.
func UnmarshalValidString(lenU mus.Unmarshaller[int], lenVl com.Validator[int],
	skip bool, bs []byte) (v string, n int, err error) {
	var length int
	if lenU == nil {
		length, n, err = varint.UnmarshalPositiveInt(bs)
	} else {
		length, n, err = lenU.UnmarshalMUS(bs)
	}
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
//
// The lenS argument specifies the Sizer for the string length.
func SizeString(v string, lenS mus.Sizer[int]) (n int) {
	return ord.SizeString(v, lenS)
}

// SkipString skips a MUS-encoded string.
//
// The lenU argument specifies the Unmarshaller for the string length.
//
// Returns the number of skiped bytes and one of the mus.ErrTooSmallByteSlice,
// com.ErrOverflow or mus.ErrNegativeLength errors.
func SkipString(lenU mus.Unmarshaller[int], bs []byte) (n int, err error) {
	return ord.SkipString(lenU, bs)
}
