package ord

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
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
func NewValidStringSerWith(lenSer mus.Serializer[int],
	lenVl com.Validator[int]) validStringSer {
	return validStringSer{NewStringSerWith(lenSer), lenVl}
}

type stringSer struct {
	lenSer mus.Serializer[int]
}

// Marshal fills bs with an encoded string value.
//
// Returns the number of used bytes. It will panic if bs is too small.
func (s stringSer) Marshal(v string, bs []byte) (n int) {
	return MarshalString(v, s.lenSer, bs)
}

// Unmarshal parses an encoded string value from bs.
//
// In addition to the string value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice, com.ErrNegativeLength, or a length
// unmarshalling error.
func (s stringSer) Unmarshal(bs []byte) (v string, n int, err error) {
	length, n, err := s.lenSer.Unmarshal(bs)
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
	return string(bs[n:l]), l, nil
}

// Size returns the size of an encoded string value.
func (s stringSer) Size(v string) (size int) {
	return SizeString(v, s.lenSer)
}

// Skip skips an encoded string value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice, mus.ErrNegativeLength, or a length unmarshalling
// error.
func (s stringSer) Skip(bs []byte) (n int, err error) {
	return SkipString(s.lenSer, bs)
}

// -----------------------------------------------------------------------------

type validStringSer struct {
	stringSer
	lenVl com.Validator[int]
}

// Unmarshal parses an encoded string value from bs.
//
// In addition to the string value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice, com.ErrNegativeLength, or a length
// unmarshalling error, or a length validation error.
func (s validStringSer) Unmarshal(bs []byte) (v string, n int, err error) {
	length, n, err := s.lenSer.Unmarshal(bs)
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
	return string(bs[n:l]), l, nil
}

// -----------------------------------------------------------------------------

func MarshalString(v string, lenSer mus.Serializer[int], bs []byte) (n int) {
	length := len(v)
	n = lenSer.Marshal(length, bs)
	if len(bs) < n+length {
		panic(mus.ErrTooSmallByteSlice)
	}
	return n + copy(bs[n:], v)
}

func SizeString(v string, lenSer mus.Serializer[int]) (size int) {
	length := len(v)
	return lenSer.Size(length) + length
}

func SkipString(lenSer mus.Serializer[int], bs []byte) (n int, err error) {
	length, n, err := lenSer.Unmarshal(bs)
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
