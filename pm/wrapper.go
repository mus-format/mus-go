package pm

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
)

type wrapper[T any] struct {
	ptrMap    *com.PtrMap
	revPtrMap *com.ReversePtrMap
	ser       mus.Serializer[T]
}

// Marshal writes an encoded pointer.
//
// In addition to the number of bytes written, it may also return an inner
// serializer marshalling error.
func (p wrapper[T]) Marshal(v T, bs []byte) (n int) {
	defer func() {
		*p.ptrMap = *com.NewPtrMap()
	}()
	return p.ser.Marshal(v, bs)
}

// Unmarshal reads an encoded pointer.
//
// In addition to the pointer and the number of bytes read, it may also return
// an inner serializer unmarshalling error.
func (p wrapper[T]) Unmarshal(bs []byte) (t T, n int, err error) {
	defer func() {
		*p.revPtrMap = *com.NewReversePtrMap()
	}()
	return p.ser.Unmarshal(bs)
}

// Size returns the size of an encoded pointer.
func (p wrapper[T]) Size(v T) int {
	defer func() {
		*p.ptrMap = *com.NewPtrMap()
	}()
	return p.ser.Size(v)
}

// Skip skips an encoded pointer.
//
// In addition to the number of bytes read, it may also return an inner
// serializer error.
func (p wrapper[T]) Skip(bs []byte) (n int, err error) {
	defer func() {
		*p.revPtrMap = *com.NewReversePtrMap()
	}()
	return p.ser.Skip(bs)
}
