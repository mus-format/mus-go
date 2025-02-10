package ord

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
)

// NewPtrMarshallerFn creates a pointer Marshaller.
func NewPtrMarshallerFn[T any](m mus.Marshaller[T]) mus.MarshallerFn[*T] {
	return func(v *T, bs []byte) (n int) {
		return MarshalPtr(v, m, bs)
	}
}

// NewPtrUnmarshallerFn creates a pointer Unmarshaller.
func NewPtrUnmarshallerFn[T any](u mus.Unmarshaller[T]) mus.UnmarshallerFn[*T] {
	return func(bs []byte) (v *T, n int, err error) {
		return UnmarshalPtr(u, bs)
	}
}

// NewPtrSizerFn creates a pointer Sizer.
func NewPtrSizerFn[T any](s mus.Sizer[T]) mus.SizerFn[*T] {
	return func(v *T) (size int) {
		return SizePtr(v, s)
	}
}

// NewPtrSkipperFn creates a pointer Skipper.
func NewPtrSkipperFn(sk mus.Skipper) mus.SkipperFn {
	return func(bs []byte) (n int, err error) {
		return SkipPtr(sk, bs)
	}
}

// MarshalPtr fills bs with the encoded pointer value.
//
// The m argument specifies the Marshaller for the pointer base type.
//
// Returns the number of used bytes. It will panic if bs is too small.
func MarshalPtr[T any](v *T, m mus.Marshaller[T], bs []byte) (n int) {
	if v == nil {
		bs[0] = byte(com.Nil)
		n = 1
		return
	}
	bs[0] = byte(com.NotNil)
	return 1 + m.Marshal(*v, bs[1:])
}

// UnmarshalPtr parses an encoded pointer value from bs.
//
// The u argument specifies the Unmarshaller for the base pointer type.
//
// In addition to the pointer value and the number of used bytes, it can
// return mus.ErrTooSmallByteSlice, com.ErrWrongFormat or Unarshaller error.
func UnmarshalPtr[T any](u mus.Unmarshaller[T], bs []byte) (v *T, n int,
	err error) {
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
	k, n, err := u.Unmarshal(bs[1:])
	if err != nil {
		n = 1 + n
		return
	}
	return &k, 1 + n, err
}

// SizePtr returns the size of an encoded pointer value.
//
// The s argument specifies the Sizer for the pointer base type.
func SizePtr[T any](v *T, s mus.Sizer[T]) (size int) {
	if v != nil {
		return 1 + s.Size(*v)
	}
	return 1
}

// SkipPtr skips an encoded pointer value.
//
// The sk argument specifies the Skipper for the pointer base type.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice, com.ErrWrongFormat or Skipper error.
func SkipPtr(sk mus.Skipper, bs []byte) (n int, err error) {
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
	n, err = sk.Skip(bs[1:])
	return 1 + n, err
}
