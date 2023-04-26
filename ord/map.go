package ord

import (
	muscom "github.com/mus-format/mus-common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/varint"
)

// MarshalMap fills bs with the MUS encoding of a map. Returns the number of
// used bytes.
func MarshalMap[T comparable, V any](v map[T]V, m1 mus.Marshaler[T],
	m2 mus.Marshaler[V],
	bs []byte,
) (n int) {
	n = varint.MarshalInt(len(v), bs)
	for k, v := range v {
		n += m1.MarshalMUS(k, bs[n:])
		n += m2.MarshalMUS(v, bs[n:])
	}
	return
}

// UnmarshalMap parses a MUS-encoded map from bs. In addition to the map, it
// returns the number of used bytes and an error.
//
// Arguments u1, u2 specify Unmarshalers for map keys and values, respectively.
//
// The error returned by UnmarshalMap can be one of mus.ErrTooSmallByteSlice,
// muscom.ErrOverflow, muscom.ErrNegativeLength, or an Unmarshaler error.
func UnmarshalMap[T comparable, V any](u1 mus.Unmarshaler[T],
	u2 mus.Unmarshaler[V],
	bs []byte,
) (v map[T]V, n int, err error) {
	return UnmarshalValidMap(nil, u1, u2, nil, nil, nil, nil, bs)
}

// UnmarshalValidMap parses a MUS-encoded valid map from bs. In addition to
// the map, it returns the number of used bytes and an error.
//
// The maxLength argument specifies the map length Validator. Arguments
// u1, u2, vl1, vl2, sk1, sk2 - Unmarshalers, Validators and Skippers for the
// map keys and values, respectively.
// If one of the Validators returns an error, UnmarshalValidMap uses the
// Skippers to skip the remaining bytes of the map. If one of the Skippers is
// nil, it immediately returns a validation error.
//
// The error returned by UnmarshalValidMap can be one of
// mus.ErrTooSmallByteSlice, muscom.ErrOverflow, muscom.ErrNegativeLength, an
// Unmarshaler, Validator, or Skipper error.
func UnmarshalValidMap[T comparable, V any](maxLength muscom.Validator[int],
	u1 mus.Unmarshaler[T],
	u2 mus.Unmarshaler[V],
	vl1 muscom.Validator[T],
	vl2 muscom.Validator[V],
	sk1, sk2 mus.Skipper,
	bs []byte,
) (v map[T]V, n int, err error) {
	length, n, err := varint.UnmarshalInt(bs)
	if err != nil {
		return
	}
	if length < 0 {
		err = muscom.ErrNegativeLength
		return
	}
	var (
		n1   int
		i    int
		err1 error
		k    T
		p    V
	)
	if maxLength != nil {
		if err = maxLength.Validate(length); err != nil {
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
				n1, err1 = sk2.SkipMUS(bs[n:])
				n += n1
				if err1 != nil {
					err = err1
					return
				}
				i++
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

// SizeMap returns the size of a MUS-encoded map.
//
// Arguments s1, s2 specify Sizers for map keys and values.
func SizeMap[T comparable, V any](v map[T]V, s1 mus.Sizer[T],
	s2 mus.Sizer[V]) (size int) {
	size += varint.SizeInt(len(v))
	for k, v := range v {
		size += s1.SizeMUS(k)
		size += s2.SizeMUS(v)
	}
	return
}

// SkipMap skips a MUS-encoded map in bs. Returns the number of skiped bytes
// and an error.
//
// Arguments sk1, sk2 specify Skippers for map keys and values, respectively.
//
// The error returned by SkipMap can be one of mus.ErrTooSmallByteSlice,
// muscom.ErrOverflow, muscom.ErrNegativeLength, or a Skipper error.
func SkipMap(sk1, sk2 mus.Skipper, bs []byte) (n int, err error) {
	length, n, err := varint.UnmarshalInt(bs)
	if err != nil {
		return
	}
	if length < 0 {
		err = muscom.ErrNegativeLength
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
