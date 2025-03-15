package unsafe

import (
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	strops "github.com/mus-format/mus-go/options/string"
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/varint"
)

// String is a string serializer.
var String = NewStringSer()

// NewStringSer returns a new string serializer. To specify a length validator,
// use NewValidStringSer instead.
func NewStringSer(ops ...strops.SetOption) stringSer {
	o := strops.Options{}
	strops.Apply(ops, &o)

	return newStringSer(o)
}

// NewStringSer returns a new valid string serializer.
func NewValidStringSer(ops ...strops.SetOption) validStringSer {
	o := strops.Options{}
	strops.Apply(ops, &o)
	return validStringSer{newStringSer(o), o.LenVl}
}

func newStringSer(o strops.Options) stringSer {
	var lenSer mus.Serializer[int] = varint.PositiveInt
	if o.LenSer != nil {
		lenSer = o.LenSer
	}
	return stringSer{lenSer}
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
