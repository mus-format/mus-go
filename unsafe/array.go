package unsafe

import (
	"reflect"
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	arrops "github.com/mus-format/mus-go/options/array"
	"github.com/mus-format/mus-go/ord"
	"github.com/mus-format/mus-go/varint"
)

// NewArraySer returns a new array serializer with the given element serializer.
// To specify a length or element validator, use NewValidArraySer instead.
//
// Panics if T is not an array type.
func NewArraySer[T, V any](elemSer mus.Serializer[V],
	ops ...arrops.SetOption[V],
) (s arraySer[T, V]) {
	var (
		o      = arrops.Options[V]{}
		t      = reflect.TypeFor[T]()
		length = t.Len()
	)
	arrops.Apply(ops, &o)
	var (
		lenSer mus.Serializer[int] = varint.PositiveInt
	)
	if o.LenSer != nil {
		lenSer = o.LenSer
	}
	return arraySer[T, V]{
		length:  length,
		elemSer: elemSer,
		lenSer:  lenSer,
		lenVl:   newLenVl(length),
	}
}

// NewValidArraySer returns a new valid array serializer with the given element
// serializer.
//
// Panics if T is not an array type.
func NewValidArraySer[T, V any](elemSer mus.Serializer[V],
	ops ...arrops.SetOption[V],
) arraySer[T, V] {
	var (
		o      = arrops.Options[V]{}
		t      = reflect.TypeFor[T]()
		length = t.Len()
	)
	arrops.Apply(ops, &o)
	var (
		lenSer mus.Serializer[int] = varint.PositiveInt
	)
	if o.LenSer != nil {
		lenSer = o.LenSer
	}
	return arraySer[T, V]{
		length:  length,
		elemSer: elemSer,
		lenSer:  lenSer,
		lenVl:   newLenVl(length),
		elemVl:  o.ElemVl,
	}
}

type arraySer[T, V any] struct {
	length  int
	elemSer mus.Serializer[V]
	lenSer  mus.Serializer[int]
	lenVl   com.Validator[int]
	elemVl  com.Validator[V]
}

// Marshal fills bs with an encoded array value.
//
// Returns the number of used bytes. It will panic if bs is too small.
func (s arraySer[T, V]) Marshal(v T, bs []byte) (n int) {
	sl := unsafe_mod.Slice((*V)(unsafe_mod.Pointer(&v)), s.length)
	return ord.MarshalSlice(sl, s.elemSer, s.lenSer, bs)
}

// Unmarshal parses an encoded array value from bs.
//
// In addition to the slice value and the number of used bytes, it may also
// return com.ErrNegativeLength, or a length/element unmarshalling error.
func (s arraySer[T, V]) Unmarshal(bs []byte) (v T, n int, err error) {
	sl, n, err := ord.UnmarshalValidSlice(bs, s.elemSer, s.lenSer, s.lenVl, s.elemVl)
	if err != nil {
		return
	}
	v = *(*T)(unsafe_mod.Pointer(unsafe_mod.SliceData(sl)))
	return
}

// Size returns the size of an encoded array value.
func (s arraySer[T, V]) Size(v T) (size int) {
	sl := unsafe_mod.Slice((*V)(unsafe_mod.Pointer(&v)), s.length)
	return ord.SizeSlice(sl, s.elemSer, s.lenSer)
}

// Skip skips an encoded array value.
//
// In addition to the number of skipped bytes, it may also return
// com.ErrNegativeLength, a length unmarshalling error, or an element skipping
// error.
func (s arraySer[T, V]) Skip(bs []byte) (n int, err error) {
	return ord.SkipSlice(bs, s.elemSer, s.lenSer)
}

func newLenVl(length int) com.ValidatorFn[int] {
	return func(t int) (err error) {
		if t > length {
			err = com.ErrTooLargeLength
		}
		return
	}
}
