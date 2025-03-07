package ord

import (
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/varint"
)

func NewArraySer[T, V any](length int, elemSer mus.Serializer[V]) arraySer[T, V] {
	return NewArraySerWith[T, V](length, varint.PositiveInt, elemSer)
}

func NewArraySerWith[T, V any](length int, lenSer mus.Serializer[int],
	elemSer mus.Serializer[V]) arraySer[T, V] {
	var (
		lenVl    = newLenVl(length)
		sliceSer = NewValidSliceSerWith[V](lenSer, elemSer, lenVl, nil)
	)
	return arraySer[T, V]{length, sliceSer}
}

func NewValidArraySer[T, V any](length int, elemSer mus.Serializer[V],
	elemVl com.Validator[V]) arraySer[T, V] {
	return NewValidArraySerWith[T, V](length, varint.PositiveInt, elemSer, elemVl)
}

func NewValidArraySerWith[T, V any](length int, lenSer mus.Serializer[int],
	elemSer mus.Serializer[V], elemVl com.Validator[V]) arraySer[T, V] {
	var (
		lenVl    = newLenVl(length)
		sliceSer = NewValidSliceSerWith[V](lenSer, elemSer, lenVl, elemVl)
	)
	return arraySer[T, V]{length, sliceSer}
}

type arraySer[T, V any] struct {
	length   int
	sliceSer validSliceSer[V]
}

func (s arraySer[T, V]) Marshal(v T, bs []byte) (n int) {
	sl := unsafe_mod.Slice((*V)(unsafe_mod.Pointer(&v)), s.length)
	return s.sliceSer.Marshal(sl, bs)
}

func (s arraySer[T, V]) Unmarshal(bs []byte) (v T, n int, err error) {
	sl, n, err := s.sliceSer.Unmarshal(bs)
	if err != nil {
		return
	}
	v = *(*T)(unsafe_mod.Pointer(unsafe_mod.SliceData(sl)))
	return
}

func (s arraySer[T, V]) Size(v T) (size int) {
	sl := unsafe_mod.Slice((*V)(unsafe_mod.Pointer(&v)), s.length)
	return s.sliceSer.Size(sl)
}

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
