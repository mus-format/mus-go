package unsafe

import (
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/ord"
)

// Bool is a bool serializer.
var Bool = boolSer{}

type boolSer struct{}

// Marshal fills bs with an encoded bool value.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func (s boolSer) Marshal(v bool, bs []byte) (n int) {
	// TODO
	*(*bool)(unsafe_mod.Pointer(&bs[0])) = v
	return 1
}

// Unmarshal parses an encoded bool value from bs.
//
// In addition to the bool value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice or com.ErrWrongFormat.
func (s boolSer) Unmarshal(bs []byte) (v bool, n int, err error) {
	if len(bs) < 1 {
		return false, 0, mus.ErrTooSmallByteSlice
	}
	if bs[0] > 1 {
		err = com.ErrWrongFormat
		return
	}
	return *(*bool)(unsafe_mod.Pointer(&bs[0])), 1, nil
}

// Size returns the size of an encoded (Raw) bool value.
func (s boolSer) Size(_ bool) (size int) {
	return 1
}

// Skip skips an encoded bool value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrWrongFormat.
func (s boolSer) Skip(bs []byte) (n int, err error) {
	return ord.SkipBool(bs)
}
