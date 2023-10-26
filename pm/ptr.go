package pm

import (
	"unsafe"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/varint"
)

// MarshalPtr fills bs with the MUS encoding of a pointer.
//
// The m argument specifies the Marshaller for the pointer base type.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalPtr[T any](v *T, m mus.Marshaller[T], mp *com.PtrMap,
	bs []byte,
) (n int) {
	if v == nil {
		bs[0] = byte(com.Nil)
		return 1
	}
	bs[0] = byte(com.Mapping)
	n = 1
	id, newOne := maptr(unsafe.Pointer(v), mp)
	n += varint.MarshalInt(id, bs[n:])
	if newOne {
		n += m.MarshalMUS(*v, bs[n:])
	}
	return
}

// UnmarshalPtr parses a MUS-encoded pointer from bs.
//
// The u argument specifies the Unmarshaller for the base pointer type.
//
// In addition to the pointer, returns the number of used bytes and one of the
// mus.ErrTooSmallByteSlice, com.ErrWrongFormat or Unarshaller errors.
func UnmarshalPtr[T any](u mus.Unmarshaller[T], mp com.ReversePtrMap,
	bs []byte,
) (v *T, n int, err error) {
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
		id, n1, err = varint.UnmarshalInt(bs[n:])
		n += n1
		if err != nil {
			return
		}
		ptr, _ := mp.Get(id)
		if ptr == nil {
			v, n1, err = unmarshalData[T](id, u, mp, bs[n:])
			n += n1
		} else {
			v = (*T)(ptr)
		}
	default:
		err = com.ErrWrongFormat
	}
	return
}

// SizePtr returns the size of a MUS-encoded pointer.
//
// The s argument specifies the Sizer for the pointer base type.
func SizePtr[T any](v *T, s mus.Sizer[T], mp *com.PtrMap) (size int) {
	size = 1
	if v != nil {
		id, newOne := maptr(unsafe.Pointer(v), mp)
		size += varint.SizeInt(id)
		if newOne {
			return size + s.SizeMUS(*v)
		} else {
			return
		}
	}
	return
}

// SkipPtr skips a MUS-encoded pointer.
//
// The sk argument specifies the Skipper for the pointer base type.
//
// Returns the number of skiped bytes and one of the mus.ErrTooSmallByteSlice,
// com.ErrWrongFormat or Skipper errors.
func SkipPtr(sk mus.Skipper, mp com.ReversePtrMap, bs []byte) (n int, err error) {
	if len(bs) < 1 {
		err = mus.ErrTooSmallByteSlice
		return
	}
	switch bs[0] {
	case byte(com.Nil):
		n = 1
	case byte(com.Mapping):
		n = 1
		var (
			id int
			n1 int
		)
		id, n1, err = varint.UnmarshalInt(bs[n:])
		n += n1
		if err != nil {
			return
		}
		_, pst := mp.Get(id)
		if !pst {
			n1, err = sk.SkipMUS(bs[n:])
			n += n1
			if err != nil {
				return
			}
			mp.Put(id, nil)
		}
	default:
		err = com.ErrWrongFormat
	}
	return
}

func unmarshalData[T any](id int, u mus.Unmarshaller[T],
	mp com.ReversePtrMap,
	bs []byte,
) (v *T, n int, err error) {
	var (
		k  T
		n1 int
	)
	mp.Put(id, unsafe.Pointer(&k))
	k, n1, err = u.UnmarshalMUS(bs)
	n += n1
	if err != nil {
		return
	}
	v = &k
	return
}

func maptr(ptr unsafe.Pointer, mp *com.PtrMap) (id int, newOne bool) {
	id, pst := mp.Get(ptr)
	if !pst {
		id = mp.Put(ptr)
		newOne = true
	}
	return
}
