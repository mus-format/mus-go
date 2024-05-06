package ord

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/varint"
)

// MarshalMap fills bs with the MUS encoding of a map value.
//
// Arguments m1, m2 specify Marshallers for the keys and map values,
// respectively.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalMap[T comparable, V any](v map[T]V, m1 mus.Marshaller[T],
	m2 mus.Marshaller[V],
	bs []byte,
) (n int) {
	n = varint.MarshalInt(len(v), bs)
	for k, v := range v {
		n += m1.MarshalMUS(k, bs[n:])
		n += m2.MarshalMUS(v, bs[n:])
	}
	return
}

// UnmarshalMap parses a MUS-encoded map value from bs.
//
// Arguments u1, u2 specify Unmarshallers for the keys and map values,
// respectively.
//
// In addition to the map value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice, com.ErrOverflow, com.ErrNegativeLength or
// Unmarshaller error.
func UnmarshalMap[T comparable, V any](u1 mus.Unmarshaller[T],
	u2 mus.Unmarshaller[V],
	bs []byte,
) (v map[T]V, n int, err error) {
	return UnmarshalValidMap(nil, u1, u2, nil, nil, nil, nil, bs)
}

// UnmarshalValidMap parses a MUS-encoded valid map value from bs.
//
// The lenVl argument specifies the map length Validator, arguments
// u1, u2, vl1, vl2, sk1, sk2 - Unmarshallers, Validators and Skippers for the
// keys and map values, respectively.
// If one of the Validators returns an error, UnmarshalValidMap uses the
// Skippers to skip the remaining bytes of the map. If one of the Skippers is
// nil, it immediately returns a validation error.
//
// In addition to the map value and the number of used bytes, it can also
// return mus.ErrTooSmallByteSlice, com.ErrOverflow, com.ErrNegativeLength,
// Unmarshaller, Validator or Skipper error.
func UnmarshalValidMap[T comparable, V any](lenVl com.Validator[int],
	u1 mus.Unmarshaller[T],
	u2 mus.Unmarshaller[V],
	vl1 com.Validator[T],
	vl2 com.Validator[V],
	sk1, sk2 mus.Skipper,
	bs []byte,
) (v map[T]V, n int, err error) {
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
		k, n1, err = u1.UnmarshalMUS(bs[n:])
		n += n1
		if err != nil {
			return
		}
		if vl1 != nil {
			if err = vl1.Validate(k); err != nil {
				if sk2 != nil {
					n1, err1 = sk2.SkipMUS(bs[n:])
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
		p, n1, err = u2.UnmarshalMUS(bs[n:])
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

// SizeMap returns the size of a MUS-encoded map value.
//
// Arguments s1, s2 specify Sizers for the keys and map values.
func SizeMap[T comparable, V any](v map[T]V, s1 mus.Sizer[T],
	s2 mus.Sizer[V]) (size int) {
	size += varint.SizeInt(len(v))
	for k, v := range v {
		size += s1.SizeMUS(k)
		size += s2.SizeMUS(v)
	}
	return
}

// SkipMap skips a MUS-encoded map value.
//
// Arguments sk1, sk2 specify Skippers for the keys and map  values,
// respectively.
//
// In addition to the number of skipped bytes, it can also return
// mus.ErrTooSmallByteSlice, com.ErrNegativeLength or Skipper error.
func SkipMap(sk1, sk2 mus.Skipper, bs []byte) (n int, err error) {
	length, n, err := varint.UnmarshalInt(bs)
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
		n1, err = sk1.SkipMUS(bs[n:])
		n += n1
		if err != nil {
			return
		}
		n1, err = sk2.SkipMUS(bs[n:])
		n += n1
		if err != nil {
			return
		}
	}
	return
}
