package pm

import (
	"unsafe"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/varint"
)

// Wrap function wraps the serializer that uses one or more pm pointer
// serializers (all created with the same pointer and reverse pointer maps), so
// it can be used like a regular serializer.
func Wrap[T any](ptrMap *com.PtrMap, revPtrMap *com.ReversePtrMap,
	ser mus.Serializer[T]) wrapper[T] {
	return wrapper[T]{ptrMap, revPtrMap, ser}
}

// NewPtrSer returns a new pointer serializer with the given pointer map,
// reverse pointer map and base type serializer.
func NewPtrSer[T any](ptrMap *com.PtrMap, revPtrMap *com.ReversePtrMap,
	baseSer mus.Serializer[T]) ptrSer[T] {
	return ptrSer[T]{ptrMap, revPtrMap, baseSer}
}

type ptrSer[T any] struct {
	ptrMap    *com.PtrMap
	revPtrMap *com.ReversePtrMap
	baseSer   mus.Serializer[T]
}

// Marshal fills bs with an encoded pointer.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s ptrSer[T]) Marshal(v *T, bs []byte) (n int) {
	if v == nil {
		bs[0] = byte(com.Nil)
		return 1
	}
	bs[0] = byte(com.Mapping)
	n = 1
	id, newOne := maptr(unsafe.Pointer(v), s.ptrMap)
	n += varint.PositiveInt.Marshal(id, bs[n:])
	if newOne {
		n += s.baseSer.Marshal(*v, bs[n:])
	}
	return
}

// Unmarshal parses an encoded pointer from bs.
//
// In addition to the pointer and the number of used bytes, it can
// return mus.ErrTooSmallByteSlice, com.ErrWrongFormat, or a base type
// unmarshalling error.
func (s ptrSer[T]) Unmarshal(bs []byte) (v *T, n int, err error) {
	if len(bs) < 1 {
		err = mus.ErrTooSmallByteSlice
		return
	}
	switch bs[0] {
	case byte(com.Nil):
		n = 1
		return
	case byte(com.Mapping):
		var (
			n1 int
			id int
		)
		n = 1
		id, n1, err = varint.PositiveInt.Unmarshal(bs[n:])
		n += n1
		if err != nil {
			return
		}
		ptr, _ := s.revPtrMap.Get(id)
		if ptr == nil {
			v, n1, err = unmarshalData[T](id, s.baseSer, s.revPtrMap, bs[n:])
			n += n1
		} else {
			v = (*T)(ptr)
		}
	default:
		err = com.ErrWrongFormat
	}
	return
}

// SizePtr returns the size of an encoded pointer.
func (s ptrSer[T]) Size(v *T) (size int) {
	size = 1
	if v != nil {
		id, newOne := maptr(unsafe.Pointer(v), s.ptrMap)
		size += varint.PositiveInt.Size(id)
		if newOne {
			return size + s.baseSer.Size(*v)
		}
	}
	return
}

// SkipPtr skips an encoded pointer.
//
// In addition to the number of skipped bytes, it can return
// mus.ErrTooSmallByteSlice, com.ErrWrongFormat, or a base type skipping error.
func (s ptrSer[T]) Skip(bs []byte) (n int, err error) {
	if len(bs) < 1 {
		err = mus.ErrTooSmallByteSlice
		return
	}
	switch bs[0] {
	case byte(com.Nil):
		n = 1
		return
	case byte(com.Mapping):
		n = 1
		var (
			id int
			n1 int
		)
		id, n1, err = varint.PositiveInt.Unmarshal(bs[n:])
		n += n1
		if err != nil {
			return
		}
		_, pst := s.revPtrMap.Get(id)
		if !pst {
			n1, err = s.baseSer.Skip(bs[n:])
			n += n1
			if err != nil {
				return
			}
			s.revPtrMap.Put(id, nil)
		}
	default:
		err = com.ErrWrongFormat
	}
	return
}

func unmarshalData[T any](id int, ser mus.Serializer[T],
	revPtrMap *com.ReversePtrMap,
	bs []byte,
) (v *T, n int, err error) {
	var (
		k  T
		n1 int
	)
	revPtrMap.Put(id, unsafe.Pointer(&k))
	k, n1, err = ser.Unmarshal(bs)
	n += n1
	if err != nil {
		return
	}
	v = &k
	return
}

func maptr(ptr unsafe.Pointer, ptrMap *com.PtrMap) (id int, newOne bool) {
	id, pst := ptrMap.Get(ptr)
	if !pst {
		id = ptrMap.Put(ptr)
		newOne = true
	}
	return
}
