package ord

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/varint"
)

// NewSliceMarshallerFn creates a slice Marshaller.
func NewSliceMarshallerFn[T any](lenM mus.Marshaller[int],
	m mus.Marshaller[T]) mus.MarshallerFn[[]T] {
	return func(v []T, bs []byte) (n int) {
		return MarshalSlice(v, lenM, m, bs)
	}
}

// NewSliceUnmarshallerFn creates a slice Unmarshaller.
func NewSliceUnmarshallerFn[T any](lenU mus.Unmarshaller[int],
	u mus.Unmarshaller[T],
) mus.UnmarshallerFn[[]T] {
	return func(bs []byte) (v []T, n int, err error) {
		return UnmarshalValidSlice(lenU, nil, u, nil, nil, bs)
	}
}

// NewValidSliceUnmarshallerFn creates a slice Unmarshaller with validation.
func NewValidSliceUnmarshallerFn[T any](lenU mus.Unmarshaller[int],
	lenVl com.Validator[int],
	u mus.Unmarshaller[T],
	vl com.Validator[T],
	sk mus.Skipper,
) mus.UnmarshallerFn[[]T] {
	return func(bs []byte) (v []T, n int, err error) {
		return UnmarshalValidSlice(lenU, lenVl, u, vl, sk, bs)
	}
}

// NewSliceSizerFn creates a slice Sizer.
func NewSliceSizerFn[T any](lenS mus.Sizer[int], s mus.Sizer[T]) mus.SizerFn[[]T] {
	return func(v []T) (size int) {
		return SizeSlice(v, lenS, s)
	}
}

// NewSliceSkipperFn creates a slice Skipper.
func NewSliceSkipperFn(lenU mus.Unmarshaller[int], sk mus.Skipper) mus.SkipperFn {
	return func(bs []byte) (n int, err error) {
		return SkipSlice(lenU, sk, bs)
	}
}

// MarshalSlice fills bs with an encoded slice value.
//
// The lenM argument specifies the Marshaller for the length of the slice, if
// nil, varint.MarshalPositiveInt() is used.
//
// The m argument specifies the Marshaller for the slice elements.
//
// Returns the number of used bytes. It will panic if bs is too small.
func MarshalSlice[T any](v []T, lenM mus.Marshaller[int], m mus.Marshaller[T],
	bs []byte) (n int) {
	if lenM == nil {
		n = varint.MarshalPositiveInt(len(v), bs)
	} else {
		n = lenM.Marshal(len(v), bs)
	}
	for _, e := range v {
		n += m.Marshal(e, bs[n:])
	}
	return
}

// UnmarshalSlice parses an encoded slice value from bs.
//
// The lenU argument specifies the Unmarshaller for the length of the slice, if
// nil, varint.UnmarshalPositiveInt() is used.
//
// The u argument specifies the Unmarshaller for the slice elements.
//
// In addition to the slice value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice, com.ErrOverflow, com.ErrNegativeLength or
// Unmarshaller error.
func UnmarshalSlice[T any](lenU mus.Unmarshaller[int], u mus.Unmarshaller[T],
	bs []byte) (v []T, n int, err error) {
	return UnmarshalValidSlice(lenU, nil, u, nil, nil, bs)
}

// UnmarshalValidSlice parses an encoded valid slice value from bs.
//
// The lenU argument specifies the Unmarshaller for the length of the slice, if
// nil, varint.UnmarshalPositiveInt() is used.
// The lenVl argument specifies the slice length Validator, arguments u,
// vl, sk - Unmarshaller, Validator and Skipper for the slice elements. If one
// of the Validators returns an error, the Skipper is used to skip the remaining
// bytes of the slice. If the Skipper is nil, it immediately returns a
// validation error.
//
// In addition to the slice value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice, com.ErrOverflow, com.ErrNegativeLength or
// Unmarshaller, Validator or Skipper error.
func UnmarshalValidSlice[T any](lenU mus.Unmarshaller[int],
	lenVl com.Validator[int],
	u mus.Unmarshaller[T],
	vl com.Validator[T],
	sk mus.Skipper,
	bs []byte,
) (v []T, n int, err error) {
	var length int
	if lenU == nil {
		length, n, err = varint.UnmarshalPositiveInt(bs)
	} else {
		length, n, err = lenU.Unmarshal(bs)
	}
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
	if lenVl != nil {
		if err = lenVl.Validate(length); err != nil {
			goto SkipRemainingBytes
		}
	}
	v = make([]T, length)
	for i = 0; i < length; i++ {
		e, n1, err = u.Unmarshal(bs[n:])
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

// SizeSlice returns the size of an encoded slice value.
//
// The lenS argument specifies the Sizer for the length of the slice, if nil,
// varint.SizePositiveInt() is used.
//
// The s argument specifies the Sizer for the slice elements.
func SizeSlice[T any](v []T, lenS mus.Sizer[int], s mus.Sizer[T]) (size int) {
	if lenS == nil {
		size = varint.SizePositiveInt(len(v))
	} else {
		size = lenS.Size(len(v))
	}
	for i := 0; i < len(v); i++ {
		size += s.Size(v[i])
	}
	return
}

// SkipSlice skips an encoded slice value.
//
// The lenU argument specifies the Unmarshaller for the length of the slice, if
// nil, varint.UnmarshalPositiveInt() is used.
// The sk argument specifies the Skipper for the slice elements.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice, com.ErrOverflow, com.ErrNegativeLength or Skipper
// error.
func SkipSlice(lenU mus.Unmarshaller[int], sk mus.Skipper, bs []byte) (n int,
	err error) {
	var length int
	if lenU == nil {
		length, n, err = varint.UnmarshalPositiveInt(bs)
	} else {
		length, n, err = lenU.Unmarshal(bs)
	}
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
		n1, err = sk.Skip(bs[n:])
		n += n1
		if err != nil {
			return
		}
	}
	return
}
