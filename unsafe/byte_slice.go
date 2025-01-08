package unsafe

import (
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/varint"
)

// MarshalByteSlice fills bs with an encoded slice value.
//
// The lenM argument specifies the Marshaller for the length of the slice, if
// nil, varint.MarshalPositiveInt() is used.
//
// Returns the number of used bytes. It will panic if bs is too small.
func MarshalByteSlice(v []byte, lenM mus.Marshaller[int], bs []byte) (n int) {
	length := len(v)
	if lenM == nil {
		n = varint.MarshalPositiveInt(length, bs)
	} else {
		n = lenM.Marshal(length, bs)
	}
	l := n + length
	if len(bs) < l {
		panic(mus.ErrTooSmallByteSlice)
	}
	return n + copy(bs[n:], unsafe_mod.Slice(unsafe_mod.SliceData(v), length))
}

// UnmarshalByteSlice parses an encoded slice value from bs.
//
// The lenU argument specifies the Unmarshaller for the length of the slice, if
// nil, varint.UnmarshalPositiveInt() is used.
//
// In addition to the slice value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice, com.ErrOverflow or com.ErrNegativeLength.
func UnmarshalByteSlice(lenU mus.Unmarshaller[int], bs []byte) (v []byte,
	n int, err error) {
	return UnmarshalValidByteSlice(lenU, nil, false, bs)
}

// UnmarshalValidSlice parses an encoded valid slice value from bs.
//
// The lenU argument specifies the Unmarshaller for the length of the slice, if
// nil, varint.UnmarshalPositiveInt() is used.
// The lenVl argument specifies the slice length Validator. If it returns
// an error and skip == true, UnmarshalValidByteSlice skips the remaining bytes
// of the slice.
//
// In addition to the slice value and the number of used bytes, it may also
// return mus.ErrTooSmallByteSlice, com.ErrOverflow or com.ErrNegativeLength.
func UnmarshalValidByteSlice(lenU mus.Unmarshaller[int],
	lenVl com.Validator[int],
	skip bool,
	bs []byte,
) (v []byte, n int, err error) {
	var length int
	if lenU == nil {
		length, n, err = varint.UnmarshalPositiveInt(bs)
	} else {
		length, n, err = lenU.Unmarshal(bs)
	}
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	l := n + length
	if len(bs) < l {
		err = mus.ErrTooSmallByteSlice
		return
	}
	if lenVl != nil {
		if err = lenVl.Validate(length); err != nil {
			if skip {
				n += length
			}
			return
		}
	}
	v = make([]byte, length)
	if length == 0 {
		return
	}
	return unsafe_mod.Slice(&bs[n], length), l, nil
}

// SizeByteSlice returns the size of an encoded slice value.
//
// The lenS argument specifies the Sizer for the length of the slice, if nil,
// varint.SizePositiveInt() is used.
func SizeByteSlice(v []byte, lenS mus.Sizer[int]) (size int) {
	length := len(v)
	if lenS == nil {
		size = varint.SizePositiveInt(length)
	} else {
		size = lenS.Size(length)
	}
	return size + length
}

// SkipByteSlice skips an encoded slice value.
//
// The lenU argument specifies the Unmarshaller for the length of the slice, if
// nil, varint.UnmarshalPositiveInt() is used.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice, com.ErrOverflow or com.ErrNegativeLength.
func SkipByteSlice(lenU mus.Unmarshaller[int], bs []byte) (n int, err error) {
	var length int
	if lenU == nil {
		length, n, err = varint.UnmarshalPositiveInt(bs)
	} else {
		length, n, err = lenU.Unmarshal(bs)
	}
	if err != nil {
		return
	}
	if length < 0 {
		err = com.ErrNegativeLength
		return
	}
	n += length
	return
}
