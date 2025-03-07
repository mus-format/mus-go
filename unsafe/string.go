package unsafe

import (
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/varint"
)

// String is a string serializer.
var String = NewStringSerWith(varint.PositiveInt)

// NewStringSerWith returns a new string serializer with the given length
// serializer.
func NewStringSerWith(lenSer mus.Serializer[int]) stringSer {
	return stringSer{lenSer}
}

// NewValidStringSer returns a new string serializer with the given length
// validator.
func NewValidStringSer(lenVl com.Validator[int]) validStringSer {
	return NewValidStringSerWith(varint.PositiveInt, lenVl)
}

// NewValidStringSerWith returns a new string serializer with the given length
// serializer and length validator.
func NewValidStringSerWith(len mus.Serializer[int],
	lenVl com.Validator[int]) validStringSer {
	return validStringSer{NewStringSerWith(len), lenVl}
}

type stringSer struct {
	len mus.Serializer[int]
}

// Marshal fills bs with an encoded string value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s stringSer) Marshal(v string, bs []byte) (n int) {
	return ord.MarshalString(v, s.len, bs)
}

// Unmarshal parses an encoded string value from bs.
//
// In addition to the string value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice, com.ErrNegativeLength, or a length
// unmarshaling error.
func (s stringSer) Unmarshal(bs []byte) (v string, n int, err error) {
	length, n, err := s.len.Unmarshal(bs)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	l := n + length
	if len(bs) < l {
		err = mus.ErrTooSmallByteSlice
		return
	}
	if length == 0 {
		return
	}
	return unsafe_mod.String(&bs[n], length), l, nil
}

// Size returns the size of an encoded string value.
func (s stringSer) Size(v string) (size int) {
	return ord.SizeString(v, s.len)
}

// Skip skips an encoded string value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice, com.ErrNegativeLength, or a length
// unmarshaling error.
func (s stringSer) Skip(bs []byte) (n int, err error) {
	return ord.SkipString(s.len, bs)
}

// -----------------------------------------------------------------------------

type validStringSer struct {
	stringSer
	lenVl com.Validator[int]
}

// Unmarshal parses an encoded string value from bs.
//
// In addition to the string value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice, com.ErrNegativeLength, a length
// unmarshaling error, or a length validation error.
func (s validStringSer) Unmarshal(bs []byte) (v string, n int, err error) {
	length, n, err := s.len.Unmarshal(bs)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	l := n + length
	if len(bs) < l {
		err = mus.ErrTooSmallByteSlice
		return
	}
	if s.lenVl != nil {
		if err = s.lenVl.Validate(length); err != nil {
			return
		}
	}
	if length == 0 {
		return
	}
	return unsafe_mod.String(&bs[n], length), l, nil
}
