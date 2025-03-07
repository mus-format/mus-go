package ord

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
)

// Bool is the bool serializer.
var Bool = boolSer{}

type boolSer struct{}

// Marshal fills bs with an encoded bool value.
//
// Returns the number of used bytes. It will panic if bs is too small.
func (s boolSer) Marshal(v bool, bs []byte) (n int) {
	if v {
		bs[0] = 1
	} else {
		bs[0] = 0
	}
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
		return false, 0, com.ErrWrongFormat
	}
	return bs[0] == 1, 1, nil
}

// Size returns the size of an encoded bool value.
func (s boolSer) Size(_ bool) (size int) {
	return 1
}

// Skip skips an encoded bool value.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice or com.ErrWrongFormat.
func (s boolSer) Skip(bs []byte) (n int, err error) {
	return SkipBool(bs)
}

func SkipBool(bs []byte) (n int, err error) {
	if len(bs) < 1 {
		return 0, mus.ErrTooSmallByteSlice
	}
	if bs[0] == 0 || bs[0] == 1 {
		return 1, nil
	}
	return 0, com.ErrWrongFormat
}
