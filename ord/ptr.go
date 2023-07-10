package ord

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
)

// MarshalPtr fills bs with the MUS encoding of a pointer. Returns the number
// of used bytes.
//
// The m argument specifies the Marshaler for the pointer base type.
//
// It will panic if receives too small bs.
func MarshalPtr[T any](v *T, m mus.Marshaler[T], bs []byte) (n int) {
	if v == nil {
		bs[0] = com.NilFlag
		n = 1
		return
	}
	bs[0] = com.NotNilFlag
	return 1 + m.MarshalMUS(*v, bs[1:])
}

// UnmarshalPtr parses a MUS-encoded pointer from bs. In addition to the
// pointer, it returns the number of used bytes and an error.
//
// The u argument specifies the Unmarshaler for the base pointer type.
//
// The error returned by UnmarshalPtr can be one of mus.ErrTooSmallByteSlice,
// com.ErrWrongFormat, or an Unarshaler error.
func UnmarshalPtr[T any](u mus.Unmarshaler[T], bs []byte) (v *T, n int,
	err error) {
	if len(bs) < 1 {
		err = mus.ErrTooSmallByteSlice
		return
	}
	if bs[0] == com.NilFlag {
		n = 1
		return
	}
	if bs[0] != com.NotNilFlag {
		err = com.ErrWrongFormat
		return
	}
	k, n, err := u.UnmarshalMUS(bs[1:])
	if err != nil {
		n = 1 + n
		return
	}
	return &k, 1 + n, err
}

// SizePtr returns the size of a MUS-encoded pointer.
//
// The s argument specifies the Sizer for the pointer base type.
func SizePtr[T any](v *T, s mus.Sizer[T]) (size int) {
	if v != nil {
		return 1 + s.SizeMUS(*v)
	}
	return 1
}

// SkipPtr skips a MUS-encoded pointer in bs. Returns the number of skiped
// bytes and an error.
//
// The sk argument specifies the Skipper for the pointer base type.
//
// The error returned by SkipPtr can be one of mus.ErrTooSmallByteSlice,
// com.ErrWrongFormat, or a Skipper error.
func SkipPtr(sk mus.Skipper, bs []byte) (n int, err error) {
	if len(bs) < 1 {
		err = mus.ErrTooSmallByteSlice
		return
	}
	if bs[0] == com.NilFlag {
		n = 1
		return
	}
	if bs[0] != com.NotNilFlag {
		err = com.ErrWrongFormat
		return
	}
	n, err = sk.SkipMUS(bs[1:])
	return 1 + n, err
}
