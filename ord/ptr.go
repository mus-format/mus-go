package ord

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
)

// NewPtrSer returns a new pointer serializer with the given base type serializer.
func NewPtrSer[T any](baseSer mus.Serializer[T]) ptrSer[T] {
	return ptrSer[T]{baseSer}
}

type ptrSer[T any] struct {
	baseSer mus.Serializer[T]
}

// Marshal fills bs with an encoded pointer value.
//
// Returns the number of used bytes. It will panic if bs is too small.
func (s ptrSer[T]) Marshal(v *T, bs []byte) (n int) {
	if v == nil {
		bs[0] = byte(com.Nil)
		n = 1
		return
	}
	bs[0] = byte(com.NotNil)
	return 1 + s.baseSer.Marshal(*v, bs[1:])
}

// Unmarshal parses an encoded pointer value from bs.
//
// In addition to the pointer value and the number of used bytes, it can
// return mus.ErrTooSmallByteSlice, com.ErrWrongFormat or a base type
// unmarshalling error.
func (s ptrSer[T]) Unmarshal(bs []byte) (v *T, n int, err error) {
	if len(bs) < 1 {
		err = mus.ErrTooSmallByteSlice
		return
	}
	if bs[0] == byte(com.Nil) {
		n = 1
		return
	}
	if bs[0] != byte(com.NotNil) {
		err = com.ErrWrongFormat
		return
	}
	k, n, err := s.baseSer.Unmarshal(bs[1:])
	if err != nil {
		n = 1 + n
		return
	}
	return &k, 1 + n, err
}

// Size returns the size of an encoded pointer value.
func (s ptrSer[T]) Size(v *T) (size int) {
	if v != nil {
		return 1 + s.baseSer.Size(*v)
	}
	return 1
}

// Skip skips an encoded pointer value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice, com.ErrWrongFormat or a base type skipping error.
func (s ptrSer[T]) Skip(bs []byte) (n int, err error) {
	if len(bs) < 1 {
		err = mus.ErrTooSmallByteSlice
		return
	}
	if bs[0] == byte(com.Nil) {
		n = 1
		return
	}
	if bs[0] != byte(com.NotNil) {
		err = com.ErrWrongFormat
		return
	}
	n, err = s.baseSer.Skip(bs[1:])
	return 1 + n, err
}
