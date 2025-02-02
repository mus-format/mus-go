package ord

import (
	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
	"github.com/mus-format/mus-go/varint"
)

// MarshalStr fills bs with an encoded string value.
//
// Provides an implementation of the Marshaller interface for the string type.
// See MarshalString for details.
func MarshalStr(v string, bs []byte) (n int) {
	return MarshalString(v, nil, bs)
}

// UnmarshalStr parses an encoded string value from bs.
//
// Provides an implementation of the Unmarshaller interface for the string type.
// See UnmarshalString for details.
func UnmarshalStr(bs []byte) (v string, n int, err error) {
	return UnmarshalValidString(nil, nil, false, bs)
}

// SizeStr returns the size of an encoded string value.
//
// Provides an implementation of the Sizer interface for the string type.
// See SizeString for details.
func SizeStr(v string) (n int) {
	return SizeString(v, nil)
}

// SkipString skips an encoded string value.
//
// Provides an implementation of the Skipper interface for the string type.
// See SkipString for details.
func SkipStr(bs []byte) (n int, err error) {
	return SkipString(nil, bs)
}

// MarshalString fills bs with an encoded string value.
//
// The lenM argument specifies the Marshaller for the string length, if nil,
// varint.MarshalPositiveInt() is used.
//
// Returns the number of used bytes. It will panic if receives too small bs.
func MarshalString(v string, lenM mus.Marshaller[int], bs []byte) (n int) {
	length := len(v)
	if lenM == nil {
		n = varint.MarshalPositiveInt(length, bs)
	} else {
		n = lenM.Marshal(length, bs)
	}
	if len(bs) < n+length {
		panic(mus.ErrTooSmallByteSlice)
	}
	return n + copy(bs[n:], v)
}

// UnmarshalString parses an encoded string value from bs.
//
// The lenU argument specifies the Unmarshaller for the string length, if nil,
// varint.UnmarshalPositiveInt() is used.
//
// In addition to the string value and the number of used bytes, it can
// return mus.ErrTooSmallByteSlice, com.ErrOverflow or com.ErrNegativeLength.
func UnmarshalString(lenU mus.Unmarshaller[int], bs []byte) (v string,
	n int, err error) {
	return UnmarshalValidString(lenU, nil, false, bs)
}

// UnmarshalValidString parses an encoded valid string value from bs.
//
// The lenU argument specifies the Unmarshaller for the string length, if nil,
// varint.UnmarshalPositiveInt() is used.
// The lenVl argument specifies the string length Validator. If it returns
// an error and skip == true, UnmarshalValidString skips the remaining bytes of
// the string.
//
// In addition to the string value and the number of used bytes, it can
// return mus.ErrTooSmallByteSlice, com.ErrOverflow, com.ErrNegativeLength or
// Validator error.
func UnmarshalValidString(lenU mus.Unmarshaller[int],
	lenVl com.Validator[int], skip bool, bs []byte) (v string, n int, err error) {
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
	if length == 0 {
		return
	}
	return string(bs[n:l]), l, nil
}

// SizeString returns the size of an encoded string value.
//
// The lenS argument specifies the Sizer for the string length, if nil,
// varint.SizePositiveInt() is used.
func SizeString(v string, lenS mus.Sizer[int]) (n int) {
	length := len(v)
	if lenS == nil {
		return varint.SizePositiveInt(length) + length
	} else {
		return lenS.Size(length) + length
	}
}

// SkipString skips an encoded string value.
//
// The lenU argument specifies the Unmarshaller for the string length, if nil,
// varint.UnmarshalPositiveInt() is used.
//
// In addition to the number of skipped bytes, it may also return
// mus.ErrTooSmallByteSlice, com.ErrOverflow or mus.ErrNegativeLength.
func SkipString(lenU mus.Unmarshaller[int], bs []byte) (n int, err error) {
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
	if len(bs[n:]) < length {
		err = mus.ErrTooSmallByteSlice
		return
	}
	return n + length, nil
}
