package unsafe

import (
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	bslops "github.com/mus-format/mus-go/options/byte_slice"
	"github.com/mus-format/mus-go/ord"
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

// // ByteSlice is a byte slice serializer.
// var ByteSlice = NewByteSliceSerWith(varint.PositiveInt)

// // NewByteSliceSerWith returns a new byte slice serializer with the given length
// // serializer.
// func NewByteSliceSerWith(lenSer mus.Serializer[int]) byteSliceSer {
// 	return byteSliceSer{lenSer}
// }

// // NewValidByteSliceSer returns a new byte slice serializer with the given length
// // validator.
// func NewValidByteSliceSer(lenVl com.Validator[int]) validByteSliceSer {
// 	return NewValidByteSliceSerWith(varint.PositiveInt, lenVl)
// }

// // NewValidByteSliceSerWith returns a new byte slice serializer with the given
// // length serializer and length validator.
// func NewValidByteSliceSerWith(lenSer mus.Serializer[int],
// 	lenVl com.Validator[int]) validByteSliceSer {
// 	return validByteSliceSer{NewByteSliceSerWith(lenSer), lenVl}
// }

type byteSliceSer struct {
	lenSer mus.Serializer[int]
}

// Marshal fills bs with an encoded byte slice value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s byteSliceSer) Marshal(v []byte, bs []byte) (n int) {
	return ord.MarshalByteSlice(v, s.lenSer, bs)
}

// Unmarshal parses an encoded byte slice value from bs.
//
// In addition to the byte slice value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice, com.ErrNegativeLength, or a length
// unmarshaling error.
func (s byteSliceSer) Unmarshal(bs []byte) (v []byte, n int, err error) {
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
	v = make([]byte, length)
	if length == 0 {
		return
	}
	return unsafe_mod.Slice(&bs[n], length), l, nil
}

// Size returns the size of an encoded byte slice value.
func (s byteSliceSer) Size(v []byte) (size int) {
	return ord.SizeByteSlice(v, s.lenSer)
}

// Skip skips an encoded byte slice value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice, mus.ErrTooSmallByteSlice, or a length unmarshaling
// error.
func (s byteSliceSer) Skip(bs []byte) (n int, err error) {
	return ord.SkipByteSlice(s.lenSer, bs)
}

// -----------------------------------------------------------------------------

type validByteSliceSer struct {
	byteSliceSer
	lenVl com.Validator[int]
}

// Unmarshal parses an encoded byte slice value from bs.
//
// In addition to the byte slice value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice, com.ErrNegativeLength, a length
// unmarshaling error, or a length validation error.
func (s validByteSliceSer) Unmarshal(bs []byte) (v []byte, n int, err error) {
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
	v = make([]byte, length)
	if length == 0 {
		return
	}
	return unsafe_mod.Slice(&bs[n], length), l, nil
}
