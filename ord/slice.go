package ord

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	slops "github.com/mus-format/mus-go/options/slice"
	"github.com/mus-format/mus-go/varint"
)

// NewSliceSer returns a new slice serializer with the given element serializer.
// To specify a length or element validator, use NewValidStringSer instead.
func NewSliceSer[T any](elemSer mus.Serializer[T], ops ...slops.SetOption[T]) (
	s sliceSer[T]) {
	o := slops.Options[T]{}
	slops.Apply(ops, &o)

	return newSliceSer(elemSer, o)
}

// NewValidSliceSer returns a new valid slice serializer.
func NewValidSliceSer[T any](elemSer mus.Serializer[T],
	ops ...slops.SetOption[T]) validSliceSer[T] {
	o := slops.Options[T]{}
	slops.Apply(ops, &o)

	var (
		lenVl  com.Validator[int]
		elemVl com.Validator[T]
	)
	if o.LenVl != nil {
		lenVl = o.LenVl
	}
	if o.ElemVl != nil {
		elemVl = o.ElemVl
	}
	return validSliceSer[T]{
		sliceSer: newSliceSer(elemSer, o),
		lenVl:    lenVl,
		elemVl:   elemVl,
	}
}

func newSliceSer[T any](elemSer mus.Serializer[T], o slops.Options[T]) (
	s sliceSer[T]) {
	var lenSer mus.Serializer[int] = varint.PositiveInt
	if o.LenSer != nil {
		lenSer = o.LenSer
	}
	return sliceSer[T]{
		elemSer: elemSer,
		lenSer:  lenSer,
	}
}

type sliceSer[T any] struct {
	lenSer  mus.Serializer[int]
	elemSer mus.Serializer[T]
}

// Marshal fills bs with an encoded slice value.
//
// Returns the number of used bytes. It will panic if bs is too small.
func (s sliceSer[T]) Marshal(v []T, bs []byte) (n int) {
	n = s.lenSer.Marshal(len(v), bs)
	for _, e := range v {
		n += s.elemSer.Marshal(e, bs[n:])
	}
	return
}

// Unmarshal parses an encoded slice value from bs.
//
// In addition to the slice value and the number of used bytes, it may also
// return com.ErrNegativeLength, or a length/element unmarshalling error.
func (s sliceSer[T]) Unmarshal(bs []byte) (v []T, n int, err error) {
	length, n, err := s.lenSer.Unmarshal(bs)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	var (
		n1 int
		e  T
	)
	v = make([]T, length)
	for i := 0; i < length; i++ {
		e, n1, err = s.elemSer.Unmarshal(bs[n:])
		n += n1
		if err != nil {
			return
		}
		v[i] = e
	}
	return
}

// Size returns the size of an encoded slice value.
func (s sliceSer[T]) Size(v []T) (size int) {
	length := len(v)
	size = s.lenSer.Size(length)
	for i := 0; i < length; i++ {
		size += s.elemSer.Size(v[i])
	}
	return
}

// Skip skips an encoded slice value.
//
// In addition to the number of skipped bytes, it may also return
// com.ErrNegativeLength, a length unmarshalling error, or an element skipping
// error.
func (s sliceSer[T]) Skip(bs []byte) (n int, err error) {
	length, n, err := s.lenSer.Unmarshal(bs)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	var n1 int
	for i := 0; i < length; i++ {
		n1, err = s.elemSer.Skip(bs[n:])
		n += n1
		if err != nil {
			return
		}
	}
	return
}

// -----------------------------------------------------------------------------

type validSliceSer[T any] struct {
	sliceSer[T]
	lenVl  com.Validator[int]
	elemVl com.Validator[T]
}

// Unmarshal parses an encoded slice value from bs.
//
// In addition to the slice value and the number of used bytes, it may also
// return com.ErrNegativeLength, a length/element unmarshalling error, or a
// length/element validation error.
func (s validSliceSer[T]) Unmarshal(bs []byte) (v []T, n int, err error) {
	length, n, err := s.lenSer.Unmarshal(bs)
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	if s.lenVl != nil {
		if err = s.lenVl.Validate(length); err != nil {
			return
		}
	}
	var (
		n1 int
		e  T
	)
	v = make([]T, length)
	for i := 0; i < length; i++ {
		e, n1, err = s.elemSer.Unmarshal(bs[n:])
		n += n1
		if err != nil {
			return
		}
		if s.elemVl != nil {
			if err = s.elemVl.Validate(e); err != nil {
				return
			}
		}
		v[i] = e
	}
	return
}
