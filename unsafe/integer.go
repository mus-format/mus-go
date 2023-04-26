package unsafe

import (
	unsafe_mod "unsafe"

	muscom "github.com/mus-format/mus-common-go"
	"github.com/mus-format/mus-go"
)

func marshalInteger64[T muscom.Integer64](t T, bs []byte) (n int) {
	*(*T)(unsafe_mod.Pointer(&bs[0])) = t
	return muscom.Num64RawSize
}

func marshalInteger32[T muscom.Integer32](t T, bs []byte) (n int) {
	*(*T)(unsafe_mod.Pointer(&bs[0])) = t
	return muscom.Num32RawSize
}

func marshalInteger16[T muscom.Integer16](t T, bs []byte) (n int) {
	*(*T)(unsafe_mod.Pointer(&bs[0])) = t

	return muscom.Num16RawSize
}

func marshalInteger8[T muscom.Integer8](t T, bs []byte) (n int) {
	*(*T)(unsafe_mod.Pointer(&bs[0])) = t
	return muscom.Num8RawSize
}

// -----------------------------------------------------------------------------
func unmarshalInteger64[T muscom.Integer64](bs []byte) (t T, n int, err error) {
	if len(bs) < muscom.Num64RawSize {
		return t, 0, mus.ErrTooSmallByteSlice
	}
	return *(*T)(unsafe_mod.Pointer(&bs[0])), muscom.Num64RawSize, nil
}

func unmarshalInteger32[T muscom.Integer32](bs []byte) (t T, n int, err error) {
	if len(bs) < muscom.Num32RawSize {
		return t, 0, mus.ErrTooSmallByteSlice
	}
	return *(*T)(unsafe_mod.Pointer(&bs[0])), muscom.Num32RawSize, nil
}

func unmarshalInteger16[T muscom.Integer16](bs []byte) (t T, n int, err error) {
	if len(bs) < muscom.Num16RawSize {
		return t, 0, mus.ErrTooSmallByteSlice
	}
	return *(*T)(unsafe_mod.Pointer(&bs[0])), muscom.Num16RawSize, nil
}

func unmarshalInteger8[T muscom.Integer8](bs []byte) (t T, n int, err error) {
	if len(bs) < muscom.Num8RawSize {
		return t, 0, mus.ErrTooSmallByteSlice
	}
	return *(*T)(unsafe_mod.Pointer(&bs[0])), muscom.Num8RawSize, nil
}
