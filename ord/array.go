package ord

import (
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	arrops "github.com/mus-format/mus-go/options/array"
	slops "github.com/mus-format/mus-go/options/slice"
)

// NewArraySer returns a new array serializer with the given array length and
// element serializer. To specify a length or element validator, use
// NewValidArraySer instead.
func NewArraySer[T, V any](length int, elemSer mus.Serializer[V],
	ops ...arrops.SetOption[V]) (s arraySer[T, V]) {
	o := arrops.Options[V]{}
	arrops.Apply(ops, &o)

	var (
		lenVl    = newLenVl(length)
		sliceSer = NewValidSliceSer[V](elemSer, slops.WithLenSer[V](o.LenSer),
			slops.WithLenValidator[V](lenVl))
	)
	return arraySer[T, V]{length, sliceSer}
}

// NewValidArraySer returns a new valid array serializer.
func NewValidArraySer[T, V any](length int, elemSer mus.Serializer[V],
	ops ...arrops.SetOption[V]) arraySer[T, V] {
	o := arrops.Options[V]{}
	arrops.Apply(ops, &o)

	var (
		lenVl    = newLenVl(length)
		sliceSer = NewValidSliceSer[V](elemSer, slops.WithLenSer[V](o.LenSer),
			slops.WithLenValidator[V](lenVl), slops.WithElemValidator(o.ElemVl))
	)
	return arraySer[T, V]{length, sliceSer}
}

type arraySer[T, V any] struct {
	length   int
	sliceSer validSliceSer[V]
}

// Marshal fills bs with an encoded array value.
//
// Returns the number of used bytes. It will panic if bs is too small.
func (s arraySer[T, V]) Marshal(v T, bs []byte) (n int) {
	sl := unsafe_mod.Slice((*V)(unsafe_mod.Pointer(&v)), s.length)
	return s.sliceSer.Marshal(sl, bs)
}

// Unmarshal parses an encoded array value from bs.
//
// In addition to the slice value and the number of used bytes, it may also
// return com.ErrNegativeLength, or a length/element unmarshalling error.
func (s arraySer[T, V]) Unmarshal(bs []byte) (v T, n int, err error) {
	sl, n, err := s.sliceSer.Unmarshal(bs)
	if err != nil {
		return
	}
	v = *(*T)(unsafe_mod.Pointer(unsafe_mod.SliceData(sl)))
	return
}

// Size returns the size of an encoded array value.
func (s arraySer[T, V]) Size(v T) (size int) {
	sl := unsafe_mod.Slice((*V)(unsafe_mod.Pointer(&v)), s.length)
	return s.sliceSer.Size(sl)
}

// Skip skips an encoded array value.
//
// In addition to the number of skipped bytes, it may also return
// com.ErrNegativeLength, a length unmarshalling error, or an element skipping
// error.
func (s arraySer[T, V]) Skip(bs []byte) (n int, err error) {
	return s.sliceSer.Skip(bs)
}

func newLenVl(length int) com.ValidatorFn[int] {
	return func(t int) (err error) {
		if t > length {
			err = com.ErrTooLargeLength
		}
		return
	}
}
