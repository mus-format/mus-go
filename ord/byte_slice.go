package ord

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/varint"
)

// ByteSlice is the byte slice serializer.
var ByteSlice = NewByteSliceSerWith(varint.PositiveInt)

// NewByteSliceSerWith returns a new byte slice serializer with the given length
// serializer.
func NewByteSliceSerWith(lenSer mus.Serializer[int]) byteSliceSer {
	return byteSliceSer{lenSer}
}

// NewValidByteSliceSer returns a new byte slice serializer with the given length
// validator.
func NewValidByteSliceSer(lenVl com.Validator[int]) validByteSliceSer {
	return NewValidByteSliceSerWith(varint.PositiveInt, lenVl)
}

// NewValidByteSliceSerWith returns a new byte slice serializer with the given
// length serializer and length validator.
func NewValidByteSliceSerWith(lenSer mus.Serializer[int],
	lenVl com.Validator[int]) validByteSliceSer {
	return validByteSliceSer{NewByteSliceSerWith(lenSer), lenVl}
}

type byteSliceSer struct {
	lenSer mus.Serializer[int]
}

// Marshal fills bs with an encoded slice value.
//
// Returns the number of used bytes. It will panic if bs is too small.
func (s byteSliceSer) Marshal(v []byte, bs []byte) (n int) {
	return MarshalByteSlice(v, s.lenSer, bs)
}

// Unmarshal parses an encoded slice value from bs.
//
// In addition to the slice value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice, com.ErrNegativeLength, or a length
// unmarshalling error.
func (s byteSliceSer) Unmarshal(bs []byte) (v []byte, n int, err error) {
	length, n, err := s.lenSer.Unmarshal(bs)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	if len(bs) < n+length {
		err = mus.ErrTooSmallByteSlice
		return
	}
	v = make([]byte, length)
	n += copy(v, bs[n:])
	return
}

// Size returns the size of an encoded slice value.
func (s byteSliceSer) Size(v []byte) (size int) {
	return SizeByteSlice(v, s.lenSer)
}

// Skip skips an encoded slice value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice, com.ErrNegativeLength, or a length unmarshaling
// error.
func (s byteSliceSer) Skip(bs []byte) (n int, err error) {
	return SkipByteSlice(s.lenSer, bs)
}

func MarshalByteSlice(v []byte, lenSer mus.Serializer[int], bs []byte) (n int) {
	length := len(v)
	n = lenSer.Marshal(length, bs)
	if len(bs) < n+length {
		panic(mus.ErrTooSmallByteSlice)
	}
	return n + copy(bs[n:], v)
}

func SizeByteSlice(v []byte, lenSer mus.Serializer[int]) (size int) {
	length := len(v)
	return lenSer.Size(length) + length
}

func SkipByteSlice(lenSer mus.Serializer[int], bs []byte) (n int, err error) {
	length, n, err := lenSer.Unmarshal(bs)
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
	n = l
	return
}

// -----------------------------------------------------------------------------

type validByteSliceSer struct {
	byteSliceSer
	lenVl com.Validator[int]
}

// Unmarshal parses an encoded slice value from bs.
//
// In addition to the slice value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice, com.ErrNegativeLength, a length
// unmarshalling error, or a length validation error.
func (s validByteSliceSer) Unmarshal(bs []byte) (v []byte, n int, err error) {
	length, n, err := s.lenSer.Unmarshal(bs)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	if len(bs) < n+length {
		err = mus.ErrTooSmallByteSlice
		return
	}
	if s.lenVl != nil {
		if err = s.lenVl.Validate(length); err != nil {
			return
		}
	}
	v = make([]byte, length)
	if length == 0 {
		return
	}
	n += copy(v, bs[n:])
	return
}
