package ord

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/varint"
)

// MarshalSlice fills bs with the MUS encoding of a slice value.
//
// The m argument specifies the Marshaller for the slice elements.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalSlice[T any](v []T, m mus.Marshaller[T], bs []byte) (n int) {
	n = varint.MarshalInt(len(v), bs)
	for _, e := range v {
		n += m.MarshalMUS(e, bs[n:])
	}
	return
}

// UnmarshalSlice parses a MUS-encoded slice value from bs.
//
// The u argument specifies the Unmarshaller for the slice elements.
//
// In addition to the slice value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice, com.ErrOverflow, com.ErrNegativeLength or
// Unmarshaller error.
func UnmarshalSlice[T any](u mus.Unmarshaller[T], bs []byte) (v []T, n int,
	err error) {
	return UnmarshalValidSlice(nil, u, nil, nil, bs)
}

// UnmarshalValidSlice parses a MUS-encoded valid slice value from bs.
//
// The maxLength argument specifies the slice length Validator, arguments u,
// vl, sk - Unmarshaller, Validator and Skipper for the slice elements. If one
// of the Validators returns an error, UnmarshalValidSlice uses the Skipper to
// skip the remaining bytes of the slice. If the value of the Skipper is nil, it
// immediately returns the validation error.
//
// In addition to the slice value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice, com.ErrOverflow, com.ErrNegativeLength or
// Unmarshaller, Validator or Skipper error.
func UnmarshalValidSlice[T any](maxLength com.Validator[int],
	u mus.Unmarshaller[T],
	vl com.Validator[T],
	sk mus.Skipper,
	bs []byte,
) (v []T, n int, err error) {
	length, n, err := varint.UnmarshalInt(bs)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	var (
		n1   int
		err1 error
		i    int
		e    T
	)
	if maxLength != nil {
		if err = maxLength.Validate(length); err != nil {
			goto SkipRemainingBytes
		}
	}
	v = make([]T, length)
	for i = 0; i < length; i++ {
		e, n1, err = u.UnmarshalMUS(bs[n:])
		n += n1
		if err != nil {
			return
		}
		if vl != nil {
			if err = vl.Validate(e); err != nil {
				i++
				goto SkipRemainingBytes
			}
		}
		v[i] = e
	}
	return
SkipRemainingBytes:
	if sk == nil {
		return
	}
	n1, err1 = skipRemainingSlice(i, length, sk, bs[n:])
	n += n1
	if err1 != nil {
		err = err1
	}
	return
}

// SizeSlice returns the size of a MUS-encoded slice value.
//
// The s argument specifies the Sizer for the slice elements.
func SizeSlice[T any](v []T, s mus.Sizer[T]) (size int) {
	size = varint.SizeInt(len(v))
	for i := 0; i < len(v); i++ {
		size += s.SizeMUS(v[i])
	}
	return
}

// SkipSlice skips a MUS-encoded slice value.
//
// The sk argument specifies the Skipper for the slice elements.
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice, com.ErrOverflow, com.ErrNegativeLength or Skipper
// error.
func SkipSlice(sk mus.Skipper, bs []byte) (n int, err error) {
	length, n, err := varint.UnmarshalInt(bs)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	n1, err := skipRemainingSlice(0, length, sk, bs[n:])
	n += n1
	return
}

func skipRemainingSlice(from int, length int, sk mus.Skipper, bs []byte) (n int,
	err error) {
	var n1 int
	for i := from; i < length; i++ {
		n1, err = sk.SkipMUS(bs[n:])
		n += n1
		if err != nil {
			return
		}
	}
	return
}
