package unsafe

import (
	unsafe_mod "unsafe"

	com "github.com/mus-format/common-go"
	"github.com/mus-format/mus-go"
)

func marshalInteger64[T com.Integer64](t T, bs []byte) (n int) {
	*(*T)(unsafe_mod.Pointer(&bs[0])) = t
	return com.Num64RawSize
}

func marshalInteger32[T com.Integer32](t T, bs []byte) (n int) {
	*(*T)(unsafe_mod.Pointer(&bs[0])) = t
	return com.Num32RawSize
}

func marshalInteger16[T com.Integer16](t T, bs []byte) (n int) {
	*(*T)(unsafe_mod.Pointer(&bs[0])) = t

	return com.Num16RawSize
}

func marshalInteger8[T com.Integer8](t T, bs []byte) (n int) {
	*(*T)(unsafe_mod.Pointer(&bs[0])) = t
	return com.Num8RawSize
}

func unmarshalInteger64[T com.Integer64](bs []byte) (t T, n int, err error) {
	if len(bs) < com.Num64RawSize {
		return t, 0, mus.ErrTooSmallByteSlice
	}
	return *(*T)(unsafe_mod.Pointer(&bs[0])), com.Num64RawSize, nil
}

func unmarshalInteger32[T com.Integer32](bs []byte) (t T, n int, err error) {
	if len(bs) < com.Num32RawSize {
		return t, 0, mus.ErrTooSmallByteSlice
	}
	return *(*T)(unsafe_mod.Pointer(&bs[0])), com.Num32RawSize, nil
}

func unmarshalInteger16[T com.Integer16](bs []byte) (t T, n int, err error) {
	if len(bs) < com.Num16RawSize {
		return t, 0, mus.ErrTooSmallByteSlice
	}
	return *(*T)(unsafe_mod.Pointer(&bs[0])), com.Num16RawSize, nil
}

func unmarshalInteger8[T com.Integer8](bs []byte) (t T, n int, err error) {
	if len(bs) < com.Num8RawSize {
		return t, 0, mus.ErrTooSmallByteSlice
	}
	return *(*T)(unsafe_mod.Pointer(&bs[0])), com.Num8RawSize, nil
}
