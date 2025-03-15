package ord

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	bslops "github.com/mus-format/mus-go/options/byte_slice"
	"github.com/mus-format/mus-go/varint"
)

// ByteSlice is the byte slice serializer.
var ByteSlice = NewByteSliceSer()

// NewByteSliceSer returns a new byte slice serializer. To specify a length
// validator, use NewValidByteSliceSer instead.
func NewByteSliceSer(ops ...bslops.SetOption) byteSliceSer {
	o := bslops.Options{}
	bslops.Apply(ops, &o)

	return newByteSliceSer(o)
}

// NewValidByteSliceSer returns a new valid byte slice serializer.
func NewValidByteSliceSer(ops ...bslops.SetOption) validByteSliceSer {
	o := bslops.Options{}
	bslops.Apply(ops, &o)

	var lenVl com.Validator[int]
	if o.LenVl != nil {
		lenVl = o.LenVl
	}
	return validByteSliceSer{newByteSliceSer(o), lenVl}
}

func newByteSliceSer(o bslops.Options) byteSliceSer {
	var lenSer mus.Serializer[int] = varint.PositiveInt
	if o.LenSer != nil {
		lenSer = o.LenSer
	}
	return byteSliceSer{lenSer}
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
