package ord

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/varint"
)

// NewMapMarshallerFn creates a map Marshaller.
func NewMapMarshallerFn[T comparable, V any](lenM mus.Marshaller[int],
	m1 mus.Marshaller[T],
	m2 mus.Marshaller[V],
) mus.MarshallerFn[map[T]V] {
	return func(v map[T]V, bs []byte) (n int) {
		return MarshalMap(v, lenM, m1, m2, bs)
	}
}

// NewMapUnmarshallerFn creates a map Unmarshaller.
func NewMapUnmarshallerFn[T comparable, V any](lenU mus.Unmarshaller[int],
	u1 mus.Unmarshaller[T],
	u2 mus.Unmarshaller[V],
) mus.UnmarshallerFn[map[T]V] {
	return func(bs []byte) (v map[T]V, n int, err error) {
		return UnmarshalValidMap(lenU, nil, u1, u2, nil, nil, nil, nil, bs)
	}
}

// NewValidMapUnmarshallerFn creates a map Unmarshaller with validation.
func NewValidMapUnmarshallerFn[T comparable, V any](lenU mus.Unmarshaller[int],
	lenVl com.Validator[int],
	u1 mus.Unmarshaller[T],
	u2 mus.Unmarshaller[V],
	vl1 com.Validator[T],
	vl2 com.Validator[V],
	sk1, sk2 mus.Skipper,
) mus.UnmarshallerFn[map[T]V] {
	return func(bs []byte) (v map[T]V, n int, err error) {
		return UnmarshalValidMap(lenU, lenVl, u1, u2, vl1, vl2, sk1, sk2, bs)
	}
}

// NewMapSizerFn creates a map Sizer.
func NewMapSizerFn[T comparable, V any](lenS mus.Sizer[int],
	s1 mus.Sizer[T],
	s2 mus.Sizer[V],
) mus.SizerFn[map[T]V] {
	return func(v map[T]V) (size int) {
		return SizeMap(v, lenS, s1, s2)
	}
}

// NewMapSkipperFn creates a map Skipper.
func NewMapSkipperFn[T comparable, V any](lenU mus.Unmarshaller[int],
	sk1 mus.Skipper,
	sk2 mus.Skipper,
) mus.SkipperFn {
	return func(bs []byte) (n int, err error) {
		return SkipMap(lenU, sk1, sk2, bs)
	}
}

// MarshalMap fills bs with the encoded map value.
//
// The lenM argument specifies the Marshaller for the length of the map, if nil,
// varint.MarshalPositiveInt() is used.
// Arguments m1, m2 specify Marshallers for the keys and values respectively.
//
// Returns the number of used bytes. It will panic if bs is too small.
func MarshalMap[T comparable, V any](v map[T]V, lenM mus.Marshaller[int],
	m1 mus.Marshaller[T],
	m2 mus.Marshaller[V],
	bs []byte,
) (n int) {
	if lenM == nil {
		n = varint.MarshalPositiveInt(len(v), bs)
	} else {
		n = lenM.Marshal(len(v), bs)
	}
	for k, v := range v {
		n += m1.Marshal(k, bs[n:])
		n += m2.Marshal(v, bs[n:])
	}
	return
}

// UnmarshalMap parses an encoded map value from bs.
//
// The lenU argument specifies the Unmarshaller for the length of the map, if
// nil, varint.UnmarshalPositiveInt() is used.
// Arguments u1, u2 specify Unmarshallers for the keys and values respectively.
//
// In addition to the map value and the number of used bytes, it may also return
// mus.ErrTooSmallByteSlice, com.ErrOverflow, com.ErrNegativeLength or
// Unmarshaller error.
func UnmarshalMap[T comparable, V any](lenU mus.Unmarshaller[int],
	u1 mus.Unmarshaller[T],
	u2 mus.Unmarshaller[V],
	bs []byte,
) (v map[T]V, n int, err error) {
	return UnmarshalValidMap(lenU, nil, u1, u2, nil, nil, nil, nil, bs)
}

// UnmarshalValidMap parses an encoded map value from bs.
//
// The lenU argument specifies the Unmarshaller for the length of the map, if
// nil, varint.UnmarshalPositiveInt() is used.
// The lenVl argument specifies the map length Validator, arguments
// u1, u2, vl1, vl2, sk1, sk2 - Unmarshallers, Validators and Skippers for the
// keys and values respectively.
// If one of the Validators returns an error, the Skippers is used to skip the
// remaining bytes of the map. If one of the Skippers is nil, it immediately
// returns a validation error.
//
// In addition to the map value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice, com.ErrOverflow, com.ErrNegativeLength,
// Unmarshaller, Validator or Skipper error.
func UnmarshalValidMap[T comparable, V any](lenU mus.Unmarshaller[int],
	lenVl com.Validator[int],
	u1 mus.Unmarshaller[T],
	u2 mus.Unmarshaller[V],
	vl1 com.Validator[T],
	vl2 com.Validator[V],
	sk1, sk2 mus.Skipper,
	bs []byte,
) (v map[T]V, n int, err error) {
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
		i    int
		err1 error
		k    T
		p    V
	)
	if lenVl != nil {
		if err = lenVl.Validate(length); err != nil {
			goto SkipRemainingBytes
		}
	}
	v = make(map[T]V)
	for i = 0; i < length; i++ {
		k, n1, err = u1.Unmarshal(bs[n:])
		n += n1
		if err != nil {
			return
		}
		if vl1 != nil {
			if err = vl1.Validate(k); err != nil {
				if sk2 != nil {
					n1, err1 = sk2.Skip(bs[n:])
					n += n1
					if err1 != nil {
						err = err1
						return
					}
					i++
				}
				goto SkipRemainingBytes
			}
		}
		p, n1, err = u2.Unmarshal(bs[n:])
		n += n1
		if err != nil {
			return
		}
		if vl2 != nil {
			if err = vl2.Validate(p); err != nil {
				i++
				goto SkipRemainingBytes
			}
		}
		v[k] = p
	}
	return
SkipRemainingBytes:
	if sk1 == nil || sk2 == nil {
		return
	}
	n1, err1 = skipRemainingMap(i, length, sk1, sk2, bs[n:])
	n += n1
	if err1 != nil {
		err = err1
	}
	return
}

// SizeMap returns the size of an encoded map value.
//
// The lenS argument specifies the Sizer for the length of the map, if nil,
// varint.SizePositiveInt() is used.
// Arguments s1, s2 specify Sizers for the keys and values.
func SizeMap[T comparable, V any](v map[T]V, lenS mus.Sizer[int],
	s1 mus.Sizer[T],
	s2 mus.Sizer[V],
) (size int) {
	if lenS == nil {
		size = varint.SizePositiveInt(len(v))
	} else {
		size = lenS.Size(len(v))
	}
	for k, v := range v {
		size += s1.Size(k)
		size += s2.Size(v)
	}
	return
}

// SkipMap skips an encoded map value.
//
// The lenU argument specifies the Unmarshaller for the length of the map, if
// nil, varint.UnmarshalPositiveInt() is used.
// Arguments sk1, sk2 specify Skippers for the keys and values respectively.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice, com.ErrNegativeLength or Skipper error.
func SkipMap(lenU mus.Unmarshaller[int], sk1, sk2 mus.Skipper, bs []byte) (
	n int, err error) {
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
	n1, err := skipRemainingMap(0, length, sk1, sk2, bs[n:])
	n += n1
	return
}

func skipRemainingMap(from int, length int, sk1, sk2 mus.Skipper, bs []byte) (
	n int, err error) {
	var n1 int
	for i := from; i < length; i++ {
		n1, err = sk1.Skip(bs[n:])
		n += n1
		if err != nil {
			return
		}
		n1, err = sk2.Skip(bs[n:])
		n += n1
		if err != nil {
			return
		}
	}
	return
}
